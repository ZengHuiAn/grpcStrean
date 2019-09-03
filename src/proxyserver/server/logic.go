package server

import (
	"sync"
)

type ServerList struct {
	g_mm map[int32]string
	lock sync.Mutex
}

var g_server_list = &ServerList{
	g_mm: make(map[int32]string),
}

func (s *ServerList) Get(srv_id int32) (string, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	srv_addr, err := s.g_mm[srv_id]
	return srv_addr, err
}

func (s *ServerList) Set(srv_id int32, srv_addr string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.g_mm[srv_id] = srv_addr
}

func (s *ServerList) Clone() map[int32]string {
	s.lock.Lock()
	defer s.lock.Unlock()

	new_mm := make(map[int32]string)
	for k, v := range s.g_mm {
		new_mm[k] = v
	}
	return new_mm
}

type sServerInfo struct {
	ID      int32  `json:"id"`
	Gateway string `json:"gateway"`
	Name    string `json:"name"`
	// Flag    int32   `json:"flag"`
}
