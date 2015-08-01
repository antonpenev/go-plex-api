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
  VideoList                     []Video     `xml:"Video"`
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
  XMLName      xml.Name   `xml:"Directory"`
  LocationList []Location `xml:"Location"`
  Count        int        `xml:"count,attr"`
  Key          string     `xml:"key,attr"`
  Title        string     `xml:"title,attr"`
  Art          string     `xml:"art,attr"`
  Composite    string     `xml:"composite,attr"`
  Filters      int        `xml:"filters,attr"`
  Refreshing   int        `xml:"refreshing,attr"`
  Thumb        string     `xml:"thumb,attr"`
  Type         string     `xml:"type,attr"`
  Agent        string     `xml:"agent,attr"`
  Scanner      string     `xml:"scanner,attr"`
  Language     string     `xml:"language,attr"`
  Uuid         string     `xml:"uuid,attr"`
  UpdatedAt    string     `xml:"updatedAt,attr"`
  CreatedAt    string     `xml:"createdAt,attr"`
  AllowSync    int        `xml:"allowSync,attr"`
}

type Location struct {
  XMLName xml.Name `xml:"Location"`
  Id      int      `xml:"id,attr"`
  Path    string   `xml:"path,attr"`
}

type Video struct {
  XMLName               xml.Name `xml:"Video"`
  RatingKey             string   `xml:"ratingKey,attr"`
  Key                   string   `xml:"key,attr"`
  Studio                string   `xml:"studio,attr"`
  Type                  string   `xml:"type,attr"`
  Title                 string   `xml:"title,attr"`
  TitleSort             string   `xml:"titleSort,attr"`
  ContentRating         string   `xml:"contentRating,attr"`
  Summary               string   `xml:"summary,attr"`
  Rating                string   `xml:"rating,attr"`
  ViewOffset            string   `xml:"viewOffset,attr"`
  LastViewedAt          string   `xml:"lastViewedAt,attr"`
  Year                  string   `xml:"year,attr"`
  Tagline               string   `xml:"tagline,attr"`
  Thumb                 string   `xml:"thumb,attr"`
  Art                   string   `xml:"art,attr"`
  Duration              string   `xml:"duration,attr"`
  OriginallyAvailableAt string   `xml:"originallyAvailableAt,attr"`
  AddedAt               string   `xml:"addedAt,attr"`
  UpdatedAt             string   `xml:"updatedAt,attr"`
  ChapterSource         string   `xml:"chapterSource,attr"`
}

type Genre struct {
  XMLName xml.Name `xml:"Genre"`
  Tag     string   `xml:"tag,attr"`
}

type Writer struct {
  XMLName xml.Name `xml:"Writer"`
  Tag     string   `xml:"tag,attr"`
}

type Country struct {
  XMLName xml.Name `xml:"Country"`
  Tag     string   `xml:"tag,attr"`
}

type Role struct {
  XMLName xml.Name `xml:"Role"`
  Tag     string   `xml:"tag,attr"`
}

type Director struct {
  XMLName xml.Name `xml:"Director"`
  Tag     string   `xml:"tag,attr"`
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

// Get all video
func (p *PlexClient) GetVideos(url string) ([]Video, error) {
  container, _ := p.fetchData(url)

  var videos []Video
  for _, video := range container.VideoList {
    videos = append(videos, video)
  }

  return videos, nil
}
