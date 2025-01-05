package api

import (
	"beatbump-server/backend/_youtube"
	"beatbump-server/backend/_youtube/api"
	"beatbump-server/backend/_youtube/api/auth"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TrendingEndpointHandler(c echo.Context) error {
	browseId := c.Param("browseId")
	if browseId == "" {
		browseId = "FEmusic_explore"
	}
	var responseBytes []byte
	var err error

	authObj := (c.(*auth.AuthContext)).AuthContext
	responseBytes, err = api.Browse(browseId, api.PageType_MusicPageTypePlaylist, "", nil, nil, nil, api.WebMusic, authObj)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error building API request: %s", err))
	}

	/*if category == "" {

		var exploreResponse _youtube.Explore
		err = json.Unmarshal(responseBytes, &exploreResponse)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error building API request: %s", err))
		}

		r := parseExplore(exploreResponse)
		return c.JSON(http.StatusOK, r)
	} else {*/
	var homeResponse _youtube.HomeResponse
	err = json.Unmarshal(responseBytes, &homeResponse)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error building API request: %s", err))
	}

	r := ParseHome(homeResponse)

	return c.JSON(http.StatusOK, r)
	//}
}
