package _youtube

import (
	"encoding/json"
	"fmt"
	"strings"
)

type RelatedResponse struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
	} `json:"responseContext"`
	Contents struct {
		SectionListRenderer struct {
			Contents []struct {
				MusicCarouselShelfRenderer struct {
					Header struct {
						MusicCarouselShelfBasicHeaderRenderer struct {
							Title struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										BrowseEndpoint      struct {
											BrowseId                              string `json:"browseId"`
											BrowseEndpointContextSupportedConfigs struct {
												BrowseEndpointContextMusicConfig struct {
													PageType string `json:"pageType"`
												} `json:"browseEndpointContextMusicConfig"`
											} `json:"browseEndpointContextSupportedConfigs"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint,omitempty"`
								} `json:"runs"`
							} `json:"title"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
							HeaderStyle    string `json:"headerStyle"`
							TrackingParams string `json:"trackingParams"`
							Strapline      struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"strapline,omitempty"`
							Thumbnail struct {
								MusicThumbnailRenderer struct {
									Thumbnail struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									ThumbnailCrop     string `json:"thumbnailCrop"`
									ThumbnailScale    string `json:"thumbnailScale"`
									TrackingParams    string `json:"trackingParams"`
									AccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibilityData"`
									OnTap struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										BrowseEndpoint      struct {
											BrowseId                              string `json:"browseId"`
											BrowseEndpointContextSupportedConfigs struct {
												BrowseEndpointContextMusicConfig struct {
													PageType string `json:"pageType"`
												} `json:"browseEndpointContextMusicConfig"`
											} `json:"browseEndpointContextSupportedConfigs"`
										} `json:"browseEndpoint"`
									} `json:"onTap"`
								} `json:"musicThumbnailRenderer"`
							} `json:"thumbnail,omitempty"`
						} `json:"musicCarouselShelfBasicHeaderRenderer"`
					} `json:"header"`
					Contents          []RelatedContents `json:"contents"`
					TrackingParams    string            `json:"trackingParams"`
					ItemSize          string            `json:"itemSize"`
					NumItemsPerColumn string            `json:"numItemsPerColumn"`
					FullbleedStyle    string            `json:"fullbleedStyle"`
				} `json:"musicCarouselShelfRenderer,omitempty"`
				MusicDescriptionShelfRenderer struct {
					Header struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"header"`
					Description struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"description"`
					MoreButton struct {
						ToggleButtonRenderer struct {
							IsToggled   bool `json:"isToggled"`
							IsDisabled  bool `json:"isDisabled"`
							DefaultIcon struct {
								IconType string `json:"iconType"`
							} `json:"defaultIcon"`
							DefaultText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"defaultText"`
							ToggledIcon struct {
								IconType string `json:"iconType"`
							} `json:"toggledIcon"`
							ToggledText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"toggledText"`
							TrackingParams string `json:"trackingParams"`
						} `json:"toggleButtonRenderer"`
					} `json:"moreButton"`
					TrackingParams string `json:"trackingParams"`
					ShelfStyle     string `json:"shelfStyle"`
				} `json:"musicDescriptionShelfRenderer,omitempty"`
			} `json:"contents"`
			TrackingParams string `json:"trackingParams"`
		} `json:"sectionListRenderer"`
	} `json:"contents"`
	TrackingParams string `json:"trackingParams"`
}

type RelatedContents struct {
	MusicResponsiveListItemRenderer *MusicResponsiveListItemRenderer `json:"musicResponsiveListItemRenderer"`
	MusicTwoRowItemRenderer         *MusicTwoRowItemRenderer         `json:"musicTwoRowItemRenderer,omitempty"`
}

func (c *RelatedContents) UnmarshalJSON(data []byte) error {

	if strings.Contains(string(data[0:80]), "musicResponsiveListItemRenderer") {
		musicResponsiveListRenderer := map[string]MusicResponsiveListItemRenderer{}
		err := json.Unmarshal(data, &musicResponsiveListRenderer)
		if err != nil {
			return err
		}
		renderer := musicResponsiveListRenderer["musicResponsiveListItemRenderer"]
		c.MusicResponsiveListItemRenderer = &renderer
		return nil
	} else if strings.Contains(string(data[0:80]), "musicTwoRowItemRenderer") {
		musicTwoRowRenderer := map[string]MusicTwoRowItemRenderer{}
		err := json.Unmarshal(data, &musicTwoRowRenderer)
		if err != nil {
			return err
		}
		renderer := musicTwoRowRenderer["musicTwoRowItemRenderer"]
		c.MusicTwoRowItemRenderer = &renderer
		return nil
	} else {
		return fmt.Errorf("unknown type: %v", string(data[0:50]))
	}
}

/*type MusicResponsiveListItemRenderer struct {
	TrackingParams string `json:"trackingParams"`
	Thumbnail      struct {
		MusicThumbnailRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			ThumbnailCrop  string `json:"thumbnailCrop"`
			ThumbnailScale string `json:"thumbnailScale"`
			TrackingParams string `json:"trackingParams"`
		} `json:"musicThumbnailRenderer"`
	} `json:"thumbnail"`
	Overlay struct {
		MusicItemThumbnailOverlayRenderer struct {
			Background struct {
				VerticalGradient struct {
					GradientLayerColors []string `json:"gradientLayerColors"`
				} `json:"verticalGradient"`
			} `json:"background"`
			Content struct {
				MusicPlayButtonRenderer struct {
					PlayNavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId                            string `json:"videoId"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint"`
					} `json:"playNavigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					PlayIcon       struct {
						IconType string `json:"iconType"`
					} `json:"playIcon"`
					PauseIcon struct {
						IconType string `json:"iconType"`
					} `json:"pauseIcon"`
					IconColor             int64 `json:"iconColor"`
					BackgroundColor       int   `json:"backgroundColor"`
					ActiveBackgroundColor int   `json:"activeBackgroundColor"`
					LoadingIndicatorColor int64 `json:"loadingIndicatorColor"`
					PlayingIcon           struct {
						IconType string `json:"iconType"`
					} `json:"playingIcon"`
					IconLoadingColor      int    `json:"iconLoadingColor"`
					ActiveScaleFactor     int    `json:"activeScaleFactor"`
					ButtonSize            string `json:"buttonSize"`
					RippleTarget          string `json:"rippleTarget"`
					AccessibilityPlayData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityPlayData"`
					AccessibilityPauseData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityPauseData"`
				} `json:"musicPlayButtonRenderer"`
			} `json:"content"`
			ContentPosition string `json:"contentPosition"`
			DisplayStyle    string `json:"displayStyle"`
		} `json:"musicItemThumbnailOverlayRenderer"`
	} `json:"overlay"`
	FlexColumns []struct {
		MusicResponsiveListItemFlexColumnRenderer struct {
			Text struct {
				Runs []struct {
					Text               string `json:"text"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId                            string `json:"videoId"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint,omitempty"`
						BrowseEndpoint struct {
							BrowseId                              string `json:"browseId"`
							BrowseEndpointContextSupportedConfigs struct {
								BrowseEndpointContextMusicConfig struct {
									PageType string `json:"pageType"`
								} `json:"browseEndpointContextMusicConfig"`
							} `json:"browseEndpointContextSupportedConfigs"`
						} `json:"browseEndpoint,omitempty"`
					} `json:"navigationEndpoint,omitempty"`
				} `json:"runs"`
			} `json:"text"`
			DisplayPriority string `json:"displayPriority"`
		} `json:"musicResponsiveListItemFlexColumnRenderer"`
	} `json:"flexColumns"`
	Menu struct {
		MenuRenderer struct {
			Items []struct {
				MenuNavigationItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId        string `json:"videoId"`
							PlaylistId     string `json:"playlistId"`
							Params         string `json:"params"`
							LoggingContext struct {
								VssLoggingContext struct {
									SerializedContextData string `json:"serializedContextData"`
								} `json:"vssLoggingContext"`
							} `json:"loggingContext"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint,omitempty"`
						ModalEndpoint struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint,omitempty"`
						BrowseEndpoint struct {
							BrowseId                              string `json:"browseId"`
							BrowseEndpointContextSupportedConfigs struct {
								BrowseEndpointContextMusicConfig struct {
									PageType string `json:"pageType"`
								} `json:"browseEndpointContextMusicConfig"`
							} `json:"browseEndpointContextSupportedConfigs"`
						} `json:"browseEndpoint,omitempty"`
						ShareEntityEndpoint struct {
							SerializedShareEntity string `json:"serializedShareEntity"`
							SharePanelType        string `json:"sharePanelType"`
						} `json:"shareEntityEndpoint,omitempty"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuNavigationItemRenderer,omitempty"`
				MenuServiceItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						QueueAddEndpoint    struct {
							QueueTarget struct {
								VideoId      string `json:"videoId"`
								OnEmptyQueue struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									WatchEndpoint       struct {
										VideoId string `json:"videoId"`
									} `json:"watchEndpoint"`
								} `json:"onEmptyQueue"`
							} `json:"queueTarget"`
							QueueInsertPosition string `json:"queueInsertPosition"`
							Commands            []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								AddToToastAction    struct {
									Item struct {
										NotificationTextRenderer struct {
											SuccessResponseText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"successResponseText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"notificationTextRenderer"`
									} `json:"item"`
								} `json:"addToToastAction"`
							} `json:"commands"`
						} `json:"queueAddEndpoint"`
					} `json:"serviceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuServiceItemRenderer,omitempty"`
				ToggleMenuServiceItemRenderer struct {
					DefaultText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"defaultText"`
					DefaultIcon struct {
						IconType string `json:"iconType"`
					} `json:"defaultIcon"`
					DefaultServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"defaultServiceEndpoint"`
					ToggledText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"toggledText"`
					ToggledIcon struct {
						IconType string `json:"iconType"`
					} `json:"toggledIcon"`
					ToggledServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						FeedbackEndpoint    struct {
							FeedbackToken string `json:"feedbackToken"`
						} `json:"feedbackEndpoint"`
					} `json:"toggledServiceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"toggleMenuServiceItemRenderer,omitempty"`
			} `json:"items"`
			TrackingParams  string `json:"trackingParams"`
			TopLevelButtons []struct {
				LikeButtonRenderer struct {
					Target struct {
						VideoId string `json:"videoId"`
					} `json:"target"`
					LikeStatus                string `json:"likeStatus"`
					TrackingParams            string `json:"trackingParams"`
					LikesAllowed              bool   `json:"likesAllowed"`
					DislikeNavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"dislikeNavigationEndpoint"`
					LikeCommand struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"likeCommand"`
				} `json:"likeButtonRenderer"`
			} `json:"topLevelButtons"`
			Accessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"menuRenderer"`
	} `json:"menu"`
	Badges []struct {
		MusicInlineBadgeRenderer struct {
			TrackingParams string `json:"trackingParams"`
			Icon           struct {
				IconType string `json:"iconType"`
			} `json:"icon"`
			AccessibilityData struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibilityData"`
		} `json:"musicInlineBadgeRenderer"`
	} `json:"badges,omitempty"`
	PlaylistItemData struct {
		VideoId string `json:"videoId"`
	} `json:"playlistItemData"`
	FlexColumnDisplayStyle string `json:"flexColumnDisplayStyle"`
	ItemHeight             string `json:"itemHeight"`
}*/

type MusicTwoRowItemRenderer struct {
	ThumbnailRenderer struct {
		MusicThumbnailRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			ThumbnailCrop  string `json:"thumbnailCrop"`
			ThumbnailScale string `json:"thumbnailScale"`
			TrackingParams string `json:"trackingParams"`
		} `json:"musicThumbnailRenderer"`
	} `json:"thumbnailRenderer"`
	AspectRatio string `json:"aspectRatio"`
	Title       struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				BrowseEndpoint      struct {
					BrowseId                              string `json:"browseId"`
					BrowseEndpointContextSupportedConfigs struct {
						BrowseEndpointContextMusicConfig struct {
							PageType string `json:"pageType"`
						} `json:"browseEndpointContextMusicConfig"`
					} `json:"browseEndpointContextSupportedConfigs"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"title"`
	Subtitle struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				BrowseEndpoint      struct {
					BrowseId                              string `json:"browseId"`
					BrowseEndpointContextSupportedConfigs struct {
						BrowseEndpointContextMusicConfig struct {
							PageType string `json:"pageType"`
						} `json:"browseEndpointContextMusicConfig"`
					} `json:"browseEndpointContextSupportedConfigs"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint,omitempty"`
		} `json:"runs"`
	} `json:"subtitle"`
	NavigationEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		WatchEndpoint       struct {
			VideoId        string `json:"videoId"`
			PlaylistId     string `json:"playlistId"`
			Params         string `json:"params"`
			LoggingContext struct {
				VssLoggingContext struct {
					SerializedContextData string `json:"serializedContextData"`
				} `json:"vssLoggingContext"`
			} `json:"loggingContext"`
			WatchEndpointMusicSupportedConfigs struct {
				WatchEndpointMusicConfig struct {
					MusicVideoType string `json:"musicVideoType"`
				} `json:"watchEndpointMusicConfig"`
			} `json:"watchEndpointMusicSupportedConfigs"`
		} `json:"watchEndpoint,omitempty"`
		BrowseEndpoint struct {
			BrowseId                              string `json:"browseId"`
			BrowseEndpointContextSupportedConfigs struct {
				BrowseEndpointContextMusicConfig struct {
					PageType string `json:"pageType"`
				} `json:"browseEndpointContextMusicConfig"`
			} `json:"browseEndpointContextSupportedConfigs"`
		} `json:"browseEndpoint"`
	} `json:"navigationEndpoint"`
	TrackingParams string `json:"trackingParams"`
	Menu           struct {
		MenuRenderer struct {
			Items []struct {
				MenuNavigationItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams   string `json:"clickTrackingParams"`
						WatchPlaylistEndpoint struct {
							PlaylistId string `json:"playlistId"`
							Params     string `json:"params"`
						} `json:"watchPlaylistEndpoint,omitempty"`
						ModalEndpoint struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint,omitempty"`
						ShareEntityEndpoint struct {
							SerializedShareEntity string `json:"serializedShareEntity"`
							SharePanelType        string `json:"sharePanelType"`
						} `json:"shareEntityEndpoint,omitempty"`
						BrowseEndpoint struct {
							BrowseId                              string `json:"browseId"`
							BrowseEndpointContextSupportedConfigs struct {
								BrowseEndpointContextMusicConfig struct {
									PageType string `json:"pageType"`
								} `json:"browseEndpointContextMusicConfig"`
							} `json:"browseEndpointContextSupportedConfigs"`
						} `json:"browseEndpoint,omitempty"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuNavigationItemRenderer,omitempty"`
				MenuServiceItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						QueueAddEndpoint    struct {
							QueueTarget struct {
								PlaylistId   string `json:"playlistId"`
								OnEmptyQueue struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									WatchEndpoint       struct {
										PlaylistId string `json:"playlistId"`
									} `json:"watchEndpoint"`
								} `json:"onEmptyQueue"`
							} `json:"queueTarget"`
							QueueInsertPosition string `json:"queueInsertPosition"`
							Commands            []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								AddToToastAction    struct {
									Item struct {
										NotificationTextRenderer struct {
											SuccessResponseText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"successResponseText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"notificationTextRenderer"`
									} `json:"item"`
								} `json:"addToToastAction"`
							} `json:"commands"`
						} `json:"queueAddEndpoint"`
					} `json:"serviceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuServiceItemRenderer,omitempty"`
				ToggleMenuServiceItemRenderer struct {
					DefaultText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"defaultText"`
					DefaultIcon struct {
						IconType string `json:"iconType"`
					} `json:"defaultIcon"`
					DefaultServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"defaultServiceEndpoint"`
					ToggledText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"toggledText"`
					ToggledIcon struct {
						IconType string `json:"iconType"`
					} `json:"toggledIcon"`
					ToggledServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						LikeEndpoint        struct {
							Status string `json:"status"`
							Target struct {
								PlaylistId string `json:"playlistId"`
							} `json:"target"`
						} `json:"likeEndpoint"`
					} `json:"toggledServiceEndpoint,omitempty"`
					TrackingParams string `json:"trackingParams"`
				} `json:"toggleMenuServiceItemRenderer,omitempty"`
			} `json:"items"`
			TrackingParams string `json:"trackingParams"`
			Accessibility  struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"menuRenderer"`
	} `json:"menu"`
	ThumbnailOverlay struct {
		MusicItemThumbnailOverlayRenderer struct {
			Background struct {
				VerticalGradient struct {
					GradientLayerColors []string `json:"gradientLayerColors"`
				} `json:"verticalGradient"`
			} `json:"background"`
			Content struct {
				MusicPlayButtonRenderer struct {
					PlayNavigationEndpoint struct {
						ClickTrackingParams   string `json:"clickTrackingParams"`
						WatchPlaylistEndpoint struct {
							PlaylistId string `json:"playlistId"`
							Params     string `json:"params,omitempty"`
						} `json:"watchPlaylistEndpoint"`
					} `json:"playNavigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					PlayIcon       struct {
						IconType string `json:"iconType"`
					} `json:"playIcon"`
					PauseIcon struct {
						IconType string `json:"iconType"`
					} `json:"pauseIcon"`
					IconColor             int64 `json:"iconColor"`
					BackgroundColor       int64 `json:"backgroundColor"`
					ActiveBackgroundColor int64 `json:"activeBackgroundColor"`
					LoadingIndicatorColor int64 `json:"loadingIndicatorColor"`
					PlayingIcon           struct {
						IconType string `json:"iconType"`
					} `json:"playingIcon"`
					IconLoadingColor      int     `json:"iconLoadingColor"`
					ActiveScaleFactor     float64 `json:"activeScaleFactor"`
					ButtonSize            string  `json:"buttonSize"`
					RippleTarget          string  `json:"rippleTarget"`
					AccessibilityPlayData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityPlayData"`
					AccessibilityPauseData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityPauseData"`
				} `json:"musicPlayButtonRenderer"`
			} `json:"content"`
			ContentPosition string `json:"contentPosition"`
			DisplayStyle    string `json:"displayStyle"`
		} `json:"musicItemThumbnailOverlayRenderer"`
	} `json:"thumbnailOverlay,omitempty"`
	ThumbnailCornerOverlay struct {
		MusicThumbnailRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			ThumbnailCrop     string `json:"thumbnailCrop"`
			TrackingParams    string `json:"trackingParams"`
			AccessibilityData struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibilityData"`
			OnTap struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				BrowseEndpoint      struct {
					BrowseId                              string `json:"browseId"`
					BrowseEndpointContextSupportedConfigs struct {
						BrowseEndpointContextMusicConfig struct {
							PageType string `json:"pageType"`
						} `json:"browseEndpointContextMusicConfig"`
					} `json:"browseEndpointContextSupportedConfigs"`
				} `json:"browseEndpoint"`
			} `json:"onTap"`
		} `json:"musicThumbnailRenderer"`
	} `json:"thumbnailCornerOverlay,omitempty"`
	SubtitleBadges []struct {
		MusicInlineBadgeRenderer struct {
			TrackingParams string `json:"trackingParams"`
			Icon           struct {
				IconType string `json:"iconType"`
			} `json:"icon"`
			AccessibilityData struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibilityData"`
		} `json:"musicInlineBadgeRenderer"`
	} `json:"subtitleBadges,omitempty"`
}
