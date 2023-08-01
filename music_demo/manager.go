package main

import "errors"

type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}
func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) FindByName(name string) (music *Music, index int) {
	if len(m.musics) == 0 {
		return nil, -1
	}
	for i, music := range m.musics {
		if music.Name == name {
			return &music, i
		}
	}
	return nil, -1
}

func (m *MusicManager) AddMusic(musics ...Music) {
	m.musics = append(m.musics, musics...)
}

func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	if index == 0 { // the first element
		m.musics = m.musics[index+1:]
	} else if index < len(m.musics)-1 { // the middle element
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else { // the last element
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *Music {
	_, index := m.FindByName(name)
	return m.Remove(index)
}
