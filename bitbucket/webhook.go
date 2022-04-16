package bitbucket

import "time"

type Webhook struct {
	Actor      string `json:"actor"`
	Repository string `json:"repository"`
	Push       struct {
		Changes []struct {
			New struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Target struct {
					Type    string    `json:"type"`
					Hash    string    `json:"hash"`
					Author  string    `json:"author"`
					Message string    `json:"message"`
					Date    time.Time `json:"date"`
					Parents []struct {
						Type  string `json:"type"`
						Hash  string `json:"hash"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"target"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"new"`
			Old struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Target struct {
					Type    string    `json:"type"`
					Hash    string    `json:"hash"`
					Author  string    `json:"author"`
					Message string    `json:"message"`
					Date    time.Time `json:"date"`
					Parents []struct {
						Type  string `json:"type"`
						Hash  string `json:"hash"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"target"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"old"`
			Links struct {
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Diff struct {
					Href string `json:"href"`
				} `json:"diff"`
				Commits struct {
					Href string `json:"href"`
				} `json:"commits"`
			} `json:"links"`
			Created bool `json:"created"`
			Forced  bool `json:"forced"`
			Closed  bool `json:"closed"`
			Commits []struct {
				Hash    string `json:"hash"`
				Type    string `json:"type"`
				Message string `json:"message"`
				Author  string `json:"author"`
				Links   struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"commits"`
			Truncated bool `json:"truncated"`
		} `json:"changes"`
	} `json:"push"`
}
