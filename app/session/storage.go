package session

import (
	"time"
	"fmt"
)

type session struct{
	UserID uint
	Dt     time.Time
}

type SessionStorage struct {
	Sessions map[string]session
}

func Init() (*SessionStorage) {

	s := &SessionStorage{
		Sessions : make(map[string]session),
	}

	s.permanentClear()

	return s
}

func (s *SessionStorage) Add(ssid string, uid uint) {

	s.Sessions[ssid] = session{
		UserID: uid,
		Dt:     time.Now(),
	}

}

func (s *SessionStorage) Kill(ssid string) {

	delete(s.Sessions, ssid)

}

func (s *SessionStorage) permanentClear() {

	go func(ss *SessionStorage) {

		for {

			for k, v := range ss.Sessions {

				if time.Since(v.Dt) >= time.Duration(3600)*time.Second {

					delete(ss.Sessions, k)

				}

			}

			fmt.Printf("Session storage cleaned, remaining %d active sessions", len(ss.Sessions))
			fmt.Println("")

			time.Sleep(time.Duration(60) * time.Second)
		}

	}(s)

}
