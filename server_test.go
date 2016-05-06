package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/arschles/assert"
	"github.com/deis/workflow-manager-api/data"
	"github.com/deis/workflow-manager/types"
	"github.com/jinzhu/gorm"
)

const (
	componentName     = "testcomponent"
	clusterID         = "testcluster"
	releaseTimeFormat = "2006-01-02T15:04:05Z"
)

func newServer(db *gorm.DB) *httptest.Server {
	// Routes consist of a path and a handler function.
	return httptest.NewServer(getRoutes(db))
}

func urlPath(ver int, remainder ...string) string {
	return fmt.Sprintf("%d/%s", ver, strings.Join(remainder, "/"))
}

// tests the GET /{apiVersion}/versions/{train}/{component}/{version} endpoint
func TestGetVersion(t *testing.T) {
	db, err := data.NewMemDB()
	assert.NoErr(t, err)
	assert.NoErr(t, data.VerifyPersistentStorage(db))
	srv := newServer(db)
	defer srv.Close()
	componentVer := types.ComponentVersion{
		Component: types.Component{Name: componentName},
		Version: types.Version{Train: "beta", Version: "2.0.0-beta-2", Released: "2016-03-31T23:54:39Z", Data: map[string]interface{}{
			"notes": "release notes",
		}},
	}
	_, err = data.SetVersion(db, componentVer)
	assert.NoErr(t, err)
	resp, err := httpGet(srv, urlPath(1, "versions", componentVer.Component.Name, componentVer.Version.Train, componentVer.Version.Version))
	assert.NoErr(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response code")
	decodedVer := new(types.ComponentVersion)
	assert.NoErr(t, json.NewDecoder(resp.Body).Decode(decodedVer))
	assert.Equal(t, *decodedVer, componentVer, "component version")
}

// tests the GET /{apiVersion}/versions/{train}/{component} endpoint
func TestGetComponentTrainVersions(t *testing.T) {
	memDB, err := data.NewMemDB()
	assert.NoErr(t, err)
	assert.NoErr(t, data.VerifyPersistentStorage(memDB))
	srv := newServer(memDB)
	defer srv.Close()
	componentVers := []types.ComponentVersion{}
	componentVer1 := types.ComponentVersion{
		Component: types.Component{Name: componentName},
		Version: types.Version{Train: "beta", Version: "2.0.0-beta-1", Released: "2016-03-30T23:54:39Z", Data: map[string]interface{}{
			"notes": "release notes",
		}},
	}
	componentVer2 := types.ComponentVersion{
		Component: types.Component{Name: componentName},
		Version: types.Version{Train: "beta", Version: "2.0.0-beta-2", Released: "2016-03-31T23:54:39Z", Data: map[string]interface{}{
			"notes": "release notes",
		}},
	}
	componentVers = append(componentVers, componentVer1)
	componentVers = append(componentVers, componentVer2)
	_, err = data.SetVersion(memDB, componentVers[0])
	assert.NoErr(t, err)
	_, err = data.SetVersion(memDB, componentVers[1])
	assert.NoErr(t, err)
	resp, err := httpGet(srv, urlPath(1, "versions", componentVer1.Component.Name, componentVer1.Version.Train))
	assert.NoErr(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response code")
	decodedVer := new([]types.ComponentVersion)
	assert.NoErr(t, json.NewDecoder(resp.Body).Decode(decodedVer))
	assert.Equal(t, *decodedVer, componentVers, "component versions")
}

// tests the GET /{apiVersion}/versions/{train}/{component}/latest endpoint
func TestGetLatestComponentTrainVersion(t *testing.T) {
	const componentName = "testcomponent"
	const train = "testtrain"

	memDB, err := data.NewMemDB()
	assert.NoErr(t, err)
	assert.NoErr(t, data.VerifyPersistentStorage(memDB))
	srv := newServer(memDB)
	defer srv.Close()

	const numCVs = 4
	const latestCVIdx = 2
	componentVersions := make([]types.ComponentVersion, numCVs)
	for i := 0; i < numCVs; i++ {
		cv := types.ComponentVersion{}
		cv.Component.Name = componentName
		cv.Component.Description = fmt.Sprintf("description%d", i)
		cv.Version.Train = train
		cv.Version.Version = fmt.Sprintf("testversion%d", i)
		cv.Version.Released = time.Now().Add(time.Duration(i) * time.Hour).Format(releaseTimeFormat)
		cv.Version.Data = map[string]interface{}{
			"notes": fmt.Sprintf("data%d", i),
		}
		if i == latestCVIdx {
			cv.Version.Released = time.Now().Add(time.Duration(numCVs+1) * time.Hour).Format(releaseTimeFormat)
		}
		if _, setErr := data.SetVersion(memDB, cv); setErr != nil {
			t.Fatalf("error setting component version %d (%s)", i, setErr)
		}
		componentVersions[i] = cv
	}

	resp, err := httpGet(srv, urlPath(2, "versions", componentName, train, "latest"))
	assert.NoErr(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response code")
	cv := new(types.ComponentVersion)
	assert.NoErr(t, json.NewDecoder(resp.Body).Decode(cv))
	exCV := componentVersions[latestCVIdx]

	assert.Equal(t, cv.Component.Name, exCV.Component.Name, "component name")
	// since the versions table doesn't store a description now, make sure it comes back empty
	assert.Equal(t, cv.Component.Description, "", "component name")

	assert.Equal(t, cv.Version.Train, exCV.Version.Train, "component version")
	assert.Equal(t, cv.Version.Version, exCV.Version.Version, "component version")
	assert.Equal(t, cv.Version.Released, exCV.Version.Released, "component release time")
	assert.Equal(t, cv.Version.Data, exCV.Version.Data, "component version data")
}

// tests the POST /{apiVersion}/versions/{train}/{component}/{version} endpoint
func TestPostVersions(t *testing.T) {
	memDB, err := data.NewMemDB()
	assert.NoErr(t, err)
	assert.NoErr(t, data.VerifyPersistentStorage(memDB))
	assert.NoErr(t, err)
	srv := newServer(memDB)
	defer srv.Close()
	train := "beta"
	version := "2.0.0-beta-2"
	componentVer := types.ComponentVersion{
		Component: types.Component{Name: componentName},
		Version: types.Version{Train: train, Version: version, Released: "2016-03-31T23:54:39Z", Data: map[string]interface{}{
			"notes": "release notes",
		}},
	}
	body := new(bytes.Buffer)
	assert.NoErr(t, json.NewEncoder(body).Encode(componentVer))
	resp, err := httpPost(srv, urlPath(2, "versions", componentName, train, version), string(body.Bytes()))
	assert.NoErr(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response code")
	retComponentVersion := new(types.ComponentVersion)
	assert.NoErr(t, json.NewDecoder(resp.Body).Decode(retComponentVersion))
	// TODO: version data property not traveling and returning as expected
	assert.Equal(t, *retComponentVersion, componentVer, "component version")
	fetchedComponentVersion, err := data.GetVersion(memDB, componentVer)
	assert.NoErr(t, err)
	assert.Equal(t, fetchedComponentVersion, componentVer, "component version")
}

// tests the GET /{apiVersion}/clusters/count endpoint
func TestGetClusters(t *testing.T) {
	memDB, err := data.NewMemDB()
	if err != nil {
		t.Fatalf("error creating new in-memory DB (%s)", err)
	}
	assert.NoErr(t, data.VerifyPersistentStorage(memDB))
	server := newServer(memDB)
	defer server.Close()
	resp, err := httpGet(server, urlPath(1, "clusters", "count"))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
}

// tests the GET /{apiVersion}/clusters/{id} endpoint
func TestGetClusterByID(t *testing.T) {
	memDB, err := data.NewMemDB()
	assert.NoErr(t, err)
	assert.NoErr(t, data.VerifyPersistentStorage(memDB))
	srv := newServer(memDB)
	defer srv.Close()
	cluster := data.ClusterStateful{}
	cluster.ID = clusterID
	cluster.Components = nil
	newCluster, err := data.CheckInAndSetCluster(memDB, clusterID, cluster)
	assert.NoErr(t, err)
	resp, err := httpGet(srv, urlPath(1, "clusters", clusterID))
	assert.NoErr(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, 200, "response code")
	decodedCluster := new(data.ClusterStateful)
	assert.NoErr(t, json.NewDecoder(resp.Body).Decode(decodedCluster))
	assert.Equal(t, *decodedCluster, newCluster, "returned cluster")
}

// tests the POST {apiVersion}/clusters/{id} endpoint
func TestPostClusters(t *testing.T) {
	memDB, err := data.NewMemDB()
	assert.NoErr(t, err)
	assert.NoErr(t, data.VerifyPersistentStorage(memDB))
	jsonData := `{"Components": [{"Component": {"Name": "component-a"}, "Version": {"Version": "1.0"}}]}`
	server := newServer(memDB)
	defer server.Close()
	resp, err := httpPost(server, urlPath(1, "clusters", clusterID), jsonData)
	if err != nil {
		t.Fatalf("POSTing to endpoint (%s)", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d", resp.StatusCode)
	}
	resp, err = httpGet(server, urlPath(1, "clusters", clusterID))
	defer resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d", resp.StatusCode)
	}
	cluster := new(data.ClusterStateful)
	if err := json.NewDecoder(resp.Body).Decode(cluster); err != nil {
		t.Fatalf("error reading response body (%s)", err)
	}
	if len(cluster.Components) <= 0 {
		t.Fatalf("no components returned")
	}
	if cluster.Components[0].Component.Name != "component-a" {
		t.Error("unexpected component name from JSON response")
	}
	// Note that we have to dereference "Version" twice because cluster.Components[0].Version
	// is itself a types.Version, which has both a "Released" and "Version" field
	if cluster.Components[0].Version.Version != "1.0" {
		t.Error("unexpected component version from JSON response")
	}
}

func httpGet(s *httptest.Server, route string) (*http.Response, error) {
	return http.Get(s.URL + "/" + route)
}

func httpPost(s *httptest.Server, route string, json string) (*http.Response, error) {
	fullURL := s.URL + "/" + route
	return http.Post(fullURL, "application/json", bytes.NewBuffer([]byte(json)))
}
