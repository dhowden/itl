// Copyright 2014, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package itl defines data types for importing iTunes Library XML (plist) files.
package itl

import (
	"io"
	"io/ioutil"
	"time"

	"github.com/dhowden/plist"
)

// Library represents the root iTunes library entity which includes a map of tracks and slice of
// playlists.
type Library struct {
	MajorVersion        int `plist:"Major Version"`
	MinorVersion        int `plist:"Minor Version"`
	Date                time.Time
	ApplicationVersion  string `plist:"Application Version"`
	Features            int
	ShowContentRatings  bool   `plist:"Show Content Ratings"`
	MusicFolder         string `plist:"Music Folder"`
	LibraryPersistentID string `plist:"Library Persistent ID"`
	Tracks              map[string]Track
	Playlists           []Playlist
}

// Track represents an iTunes library track, which is a media file which can either be music or video.
// Items are identified in iTunes using the TrackID.
type Track struct {
	TrackID int `plist:"Track ID"`
	Name    string
	Artist  string

	Composer string
	Year     int
	Genre    string
	Kind     string
	Size     int

	BPM int

	TrackNumber int `plist:"Track Number"`
	TrackCount  int `plist:"Track Count"`
	DiscNumber  int `plist:"Disc Number"`
	DiscCount   int `plist:"Disc Count"`

	PartOfGaplessAlbum bool   `plist:"Part Of Gapless Album"`
	ContentRating      string `plist:"Content Rating"`

	Rating         int
	RatingComputed bool `plist:"Rating Computed"`
	Disabled       bool
	Loved          bool `plist:"Loved"`

	Album               string
	AlbumArtist         string `plist:"Album Artist"`
	AlbumRating         int    `plist:"Album Rating"`
	AlbumRatingComputed bool   `plist:"Album Rating Computed"`
	AlbumLoved          bool   `plist:"Album Loved"`

	SortName        string `plist:"Sort Name"`
	SortArtist      string `plist:"Sort Artist"`
	SortAlbumArtist string `plist:"Sort Album Artist"`
	SortAlbum       string `plist:"Sort Album"`
	SortComposer    string `plist:"Sort Composer"`

	Clean  bool
	Series string

	TotalTime        int       `plist:"Total Time"`
	DateModified     time.Time `plist:"Date Modified"`
	DateAdded        time.Time `plist:"Date Added"`
	BitRate          int       `plist:"Bit Rate"`
	SampleRate       int       `plist:"Sample Rate"`
	VolumeAdjustment int       `plist:"Volume Adjustment"`
	Comments         string

	PlayCount   int       `plist:"Play Count"`
	PlayDate    int       `plist:"Play Date"`
	PlayDateUTC time.Time `plist:"Play Date UTC"`

	Protected bool
	Purchased bool

	SkipCount int       `plist:"Skip Count"`
	SkipDate  time.Time `plist:"Skip Date"`

	ArtworkCount int `plist:"Artwork Count"`

	Episode      string
	EpisodeOrder int  `plist:"Episode Order"`
	TVShow       bool `plist:"TV Show"`
	Season       int
	Podcast      bool
	ITunesU      bool `plist:"iTunesU"`
	Unplayed     bool

	PersistentID string `plist:"Persistent ID"`
	TrackType    string `plist:"Track Type"`
	Location     string
	FileType     int `plist:"File Type"`
	Movie        bool
	MusicVideo   bool `plist:"Music Video"`
	HD           bool
	HasVideo     bool `plist:"Has Video"`
	VideoHeight  int  `plist:"Video Height"`
	VideoWidth   int  `plist:"Video Width"`

	Grouping    string
	Compilation bool
	ReleaseDate time.Time `plist:"Release Date"`

	FileFolderCount    int `plist:"File Folder Count"`
	LibraryFolderCount int `plist:"Library Folder Count"`
}

// Playlist represents an iTunes playlist.
type Playlist struct {
	Name                 string
	Master               bool
	PlaylistID           int    `plist:"Playlist ID"`
	ParentPersistentID   string `plist:"Parent Persistent ID"`
	PlaylistPersistentID string `plist:"Playlist Persistent ID"`
	DistinguishedKind    int    `plist:"Distinguished Kind"`
	Visible              bool
	Music                bool
	Movies               bool
	TVShows              bool `plist:"TV Shows"`
	Podcasts             bool
	ITunesU              bool `plist:"iTunesU"`
	Audiobooks           bool
	AllItems             bool `plist:"All Items"`
	Folder               bool
	PlaylistItems        []PlaylistItem `plist:"Playlist Items"`
}

// PlaylistItem represents an individual track in a an iTunes playlist.
type PlaylistItem struct {
	TrackID int `plist:"Track ID"`
}

// ReadFromXML reads iTunes XML (plist) data from the underlying io.Reader
// returning the resuling Library.
func ReadFromXML(r io.Reader) (l Library, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = plist.Unmarshal(b, &l)
	return
}
