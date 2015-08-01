package plex

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

type MediaContainer struct {
  XMLName                       xml.Name    `xml:"MediaContainer"`
  DirectoryList                 []Directory `xml:"Directory"`
  Size                          int         `xml:"size,attr"`
  AllowCameraUpload             int         `xml:"allowCameraUpload,attr"`
  AllowSync                     int         `xml:"allowSync,attr"`
  AllowChannelAccess            int         `xml:"allowChannelAccess,attr"`
  RequestParametersInCookie     int         `xml:"requestParametersInCookie,attr"`
  Sync                          int         `xml:"sync,attr"`
  TranscoderActiveVideoSessions int         `xml:"transcoderActiveVideoSessions, attr`
  TranscoderAudio               int         `xml:"transcoderAudio,attr"`
  TranscoderVideo               int         `xml:"transcoderVideo,attr"`
  TranscoderVideoBitrates       string      `xml:"transcoderVideoBitrates,attr"`
  TranscoderVideoQualities      string      `xml:"transcoderVideoQualities,attr"`
  TranscoderVideoResolutions    string      `xml:"transcoderVideoResolutions,attr"`
  FriendlyName                  string      `xml:"friendlyName,attr"`
  MachineIdentifier             string      `xml:"machineIdentifier,attr"`
}

type Directory struct {
  XMLName xml.Name `xml:"Directory"`
  Count   string   `xml:"count,attr"`
  Key     string   `xml:"key,attr"`
  Title   string   `xml:"title,attr"`
}

type PlexClient struct {
  SERVER_URL string
}

// Creates new client to the given Plex Server Url
func New(server string) PlexClient {
  fmt.Println("Creating new Plex API Client...")
  cl := PlexClient{server}
  return cl
}

// Call the Plex Api to a desired endpoint/resource
func (p *PlexClient) fetchData(url string) (MediaContainer, error) {

  response, err := http.Get(p.SERVER_URL + url)
  if err != nil {
    return MediaContainer{}, err
  }

  defer response.Body.Close()
  data, _ := ioutil.ReadAll(response.Body)

  var container MediaContainer
  xml.Unmarshal(data, &container)

  return container, nil
}

// Makes call for Plex directories and return the the result
func (p *PlexClient) GetDirectories(url string) ([]Directory, error) {
  container, _ := p.fetchData(url)

  var directories []Directory
  for _, directory := range container.DirectoryList {
    directories = append(directories, directory)
  }

  return directories, nil
}
