package main

import "testing"

func TestNewMusicManager(t *testing.T) {
	manager := NewMusicManager()
	if manager == nil {
		t.Error("NewMusicManager failed.")
	}
}

func TestMusicManager_Len(t *testing.T) {
	manager := NewMusicManager()
	if manager.Len() != 0 {
		t.Error("Initialize MusicManager failed, not empty")
	}
}

func TestMusicManager_AddMusic(t *testing.T) {
	manager := NewMusicManager()
	music := &Music{
		Id:     "1",
		Name:   "勇次",
		Artist: "长渕刚",
		Source: "https://www.bilibili.com/video/BV1j4411j7Wx/",
		Type:   "MP3",
	}
	manager.AddMusic(*music)
	if manager.Len() != 1 {
		t.Error("MusicManager length error, expect one")
	}
}

func TestMusicManager_Get(t *testing.T) {
	manager := NewMusicManager()
	music := &Music{
		Id:     "1",
		Name:   "勇次",
		Artist: "长渕刚",
		Source: "https://www.bilibili.com/video/BV1j4411j7Wx/",
		Type:   "MP3",
	}
	manager.AddMusic(*music)
	m, err := manager.Get(0)
	if err != nil {
		t.Error("Get music failed")
	}
	if m.Id != "1" && m.Name != "勇次" && m.Artist != "长渕刚" {
		t.Error("Music info error.")
	}
}

func TestMusicManager_FindByName(t *testing.T) {
	manager := NewMusicManager()
	music := &Music{
		Id:     "1",
		Name:   "勇次",
		Artist: "长渕刚",
		Source: "https://www.bilibili.com/video/BV1j4411j7Wx/",
		Type:   "MP3",
	}
	manager.AddMusic(*music)
	m, _ := manager.FindByName("勇次")
	if m == nil {
		t.Error("FindByName failed")
	}
	if m.Name != "勇次" {
		t.Error("FindByName info error.")
	}
}

func TestMusicManager_Remove(t *testing.T) {
	manager := NewMusicManager()
	music := &Music{
		Id:     "1",
		Name:   "勇次",
		Artist: "长渕刚",
		Source: "https://www.bilibili.com/video/BV1j4411j7Wx/",
		Type:   "MP3",
	}
	music1 := &Music{
		Id:     "2",
		Name:   "干杯",
		Artist: "长渕刚",
		Source: "https://www.bilibili.com/video/BV1j4411j7Wx/",
		Type:   "MP3",
	}
	manager.AddMusic(*music, *music1)
	m := manager.Remove(0)
	if m == nil {
		t.Error("Remove failed")
	}
	if m.Name != "勇次" {
		t.Error("Remove info error.")
	}
	if manager.Len() != 1 {
		t.Error("MusicManager length error after remove")
	}
}
