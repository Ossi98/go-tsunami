//Exemple of struct generated from tsunami json output : https://mholt.github.io/json-to-go/
type AutoGenerated struct {
	ScanStatus   string `json:"scanStatus"`
	ScanFindings []struct {
		TargetInfo struct {
			NetworkEndpoints []struct {
				Type      string `json:"type"`
				IPAddress struct {
					AddressFamily string `json:"addressFamily"`
					Address       string `json:"address"`
				} `json:"ipAddress"`
			} `json:"networkEndpoints"`
		} `json:"targetInfo"`
		NetworkService struct {
			NetworkEndpoint struct {
				Type      string `json:"type"`
				IPAddress struct {
					AddressFamily string `json:"addressFamily"`
					Address       string `json:"address"`
				} `json:"ipAddress"`
				Port struct {
					PortNumber int `json:"portNumber"`
				} `json:"port"`
			} `json:"networkEndpoint"`
			TransportProtocol string `json:"transportProtocol"`
			ServiceName       string `json:"serviceName"`
			Software          struct {
				Name string `json:"name"`
			} `json:"software"`
			VersionSet struct {
				Versions []struct {
					Type              string `json:"type"`
					FullVersionString string `json:"fullVersionString"`
				} `json:"versions"`
			} `json:"versionSet"`
			ServiceContext struct {
				WebServiceContext struct {
					Software struct {
						Name string `json:"name"`
					} `json:"software"`
					VersionSet struct {
						Versions []struct {
							Type              string `json:"type"`
							FullVersionString string `json:"fullVersionString"`
						} `json:"versions"`
					} `json:"versionSet"`
					CrawlResults []struct {
						CrawlTarget struct {
							URL        string `json:"url"`
							HTTPMethod string `json:"httpMethod"`
						} `json:"crawlTarget"`
						CrawlDepth      int    `json:"crawlDepth,omitempty"`
						ResponseCode    int    `json:"responseCode"`
						ContentType     string `json:"contentType"`
						Content         string `json:"content,omitempty"`
						ResponseHeaders []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"responseHeaders"`
					} `json:"crawlResults"`
				} `json:"webServiceContext"`
			} `json:"serviceContext"`
			Cpes []string `json:"cpes"`
		} `json:"networkService"`
		Vulnerability struct {
			MainID struct {
				Publisher string `json:"publisher"`
				Value     string `json:"value"`
			} `json:"mainId"`
			Severity       string `json:"severity"`
			Title          string `json:"title"`
			Description    string `json:"description"`
			Recommendation string `json:"recommendation"`
		} `json:"vulnerability"`
	} `json:"scanFindings"`
	ScanStartTimestamp   time.Time `json:"scanStartTimestamp"`
	ScanDuration         string    `json:"scanDuration"`
	FullDetectionReports struct {
		DetectionReports []struct {
			TargetInfo struct {
				NetworkEndpoints []struct {
					Type      string `json:"type"`
					IPAddress struct {
						AddressFamily string `json:"addressFamily"`
						Address       string `json:"address"`
					} `json:"ipAddress"`
				} `json:"networkEndpoints"`
			} `json:"targetInfo"`
			NetworkService struct {
				NetworkEndpoint struct {
					Type      string `json:"type"`
					IPAddress struct {
						AddressFamily string `json:"addressFamily"`
						Address       string `json:"address"`
					} `json:"ipAddress"`
					Port struct {
						PortNumber int `json:"portNumber"`
					} `json:"port"`
				} `json:"networkEndpoint"`
				TransportProtocol string `json:"transportProtocol"`
				ServiceName       string `json:"serviceName"`
				Software          struct {
					Name string `json:"name"`
				} `json:"software"`
				VersionSet struct {
					Versions []struct {
						Type              string `json:"type"`
						FullVersionString string `json:"fullVersionString"`
					} `json:"versions"`
				} `json:"versionSet"`
				ServiceContext struct {
					WebServiceContext struct {
						Software struct {
							Name string `json:"name"`
						} `json:"software"`
						VersionSet struct {
							Versions []struct {
								Type              string `json:"type"`
								FullVersionString string `json:"fullVersionString"`
							} `json:"versions"`
						} `json:"versionSet"`
						CrawlResults []struct {
							CrawlTarget struct {
								URL        string `json:"url"`
								HTTPMethod string `json:"httpMethod"`
							} `json:"crawlTarget"`
							CrawlDepth      int    `json:"crawlDepth,omitempty"`
							ResponseCode    int    `json:"responseCode"`
							ContentType     string `json:"contentType"`
							Content         string `json:"content,omitempty"`
							ResponseHeaders []struct {
								Key   string `json:"key"`
								Value string `json:"value"`
							} `json:"responseHeaders"`
						} `json:"crawlResults"`
					} `json:"webServiceContext"`
				} `json:"serviceContext"`
				Cpes []string `json:"cpes"`
			} `json:"networkService"`
			DetectionTimestamp time.Time `json:"detectionTimestamp"`
			DetectionStatus    string    `json:"detectionStatus"`
			Vulnerability      struct {
				MainID struct {
					Publisher string `json:"publisher"`
					Value     string `json:"value"`
				} `json:"mainId"`
				Severity       string `json:"severity"`
				Title          string `json:"title"`
				Description    string `json:"description"`
				Recommendation string `json:"recommendation"`
			} `json:"vulnerability"`
		} `json:"detectionReports"`
	} `json:"fullDetectionReports"`
	ReconnaissanceReport struct {
		TargetInfo struct {
			NetworkEndpoints []struct {
				Type      string `json:"type"`
				IPAddress struct {
					AddressFamily string `json:"addressFamily"`
					Address       string `json:"address"`
				} `json:"ipAddress"`
			} `json:"networkEndpoints"`
		} `json:"targetInfo"`
		NetworkServices []struct {
			NetworkEndpoint struct {
				Type      string `json:"type"`
				IPAddress struct {
					AddressFamily string `json:"addressFamily"`
					Address       string `json:"address"`
				} `json:"ipAddress"`
				Port struct {
					PortNumber int `json:"portNumber"`
				} `json:"port"`
			} `json:"networkEndpoint"`
			TransportProtocol string `json:"transportProtocol"`
			ServiceName       string `json:"serviceName"`
			Software          struct {
				Name string `json:"name"`
			} `json:"software,omitempty"`
			VersionSet struct {
				Versions []struct {
					Type              string `json:"type"`
					FullVersionString string `json:"fullVersionString"`
				} `json:"versions"`
			} `json:"versionSet,omitempty"`
			ServiceContext struct {
				WebServiceContext struct {
					Software struct {
						Name string `json:"name"`
					} `json:"software"`
					VersionSet struct {
						Versions []struct {
							Type              string `json:"type"`
							FullVersionString string `json:"fullVersionString"`
						} `json:"versions"`
					} `json:"versionSet"`
					CrawlResults []struct {
						CrawlTarget struct {
							URL        string `json:"url"`
							HTTPMethod string `json:"httpMethod"`
						} `json:"crawlTarget"`
						CrawlDepth      int    `json:"crawlDepth,omitempty"`
						ResponseCode    int    `json:"responseCode"`
						ContentType     string `json:"contentType"`
						Content         string `json:"content,omitempty"`
						ResponseHeaders []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"responseHeaders"`
					} `json:"crawlResults"`
				} `json:"webServiceContext"`
			} `json:"serviceContext,omitempty"`
			Cpes []string `json:"cpes,omitempty"`
		} `json:"networkServices"`
	} `json:"reconnaissanceReport"`
}
