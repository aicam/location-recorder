package internal

func (s *Server) Route() {
	s.Router.POST("/recordloc", s.RecordLoc())
}
