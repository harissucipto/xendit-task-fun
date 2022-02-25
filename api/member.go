package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type ListMemberRequestURI struct {
	OrgName string `uri:"org-name" binding:"required"`
}

type ListMembersResponse []struct {
   Login string    `json:"login"`
}

// get list of member from  https://api.github.com/orgs/github/members
func (server *Server) listMembers(ctx *gin.Context) {

	var requestURI ListMemberRequestURI

	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.CheckIsValidOrg(requestURI.OrgName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	members, err := server.GetListMembers(requestURI.OrgName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	membersWithFollowers, err := server.GetNumberOfFollowerFromListMember(members)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// sort the members with followers with descending order for number of followers
	sort.Slice(membersWithFollowers, func(i, j int) bool {
		return membersWithFollowers[i].Followers > membersWithFollowers[j].Followers
	})

	
	ctx.JSON(http.StatusOK, membersWithFollowers)
}

// function to check if github organiztion is valid from api
func (server *Server) CheckIsValidOrg (org string) (bool, error) {

	url := fmt.Sprintf("%s/orgs/%s", server.config.GithubEndpoint, org)
	fmt.Println(url,"url")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header =  http.Header{
    "Authorization": []string{fmt.Sprintf("Bearer %s", server.config.GithubToken)},
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	// if resp status not 200 then retun error
	if resp.StatusCode != http.StatusOK {
		// resp status not found then return eror
		if resp.StatusCode == http.StatusNotFound {
			return false, fmt.Errorf("org not found")
		}
		return false, fmt.Errorf("github api error")
	}
	
	defer resp.Body.Close()

	return true, nil
}

type ListMemberOrg []struct {
   Login string    `json:"login"`
	 AvatarURL string `json:"avatar_url"`
	 URL string `json:"url"`
}

// function to get list of member from  https://api.github.com/orgs/github/members
func (server *Server) GetListMembers (org string) (ListMemberOrg, error) {
	url := fmt.Sprintf("%s/orgs/%s/members", server.config.GithubEndpoint,  org)
	fmt.Println(url,"url")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header =  http.Header{
    "Authorization": []string{fmt.Sprintf("Bearer %s", server.config.GithubToken)},
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// if resp status not 200 then retun error
	if resp.StatusCode != http.StatusOK	{
		// resp status not found then return eror
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("org not found")
		}
		return nil, fmt.Errorf("github api error")
	}

	defer resp.Body.Close()

	//Create a variable of the same type as our model
	 var cResp ListMemberOrg

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return nil, err
	}

	return cResp, nil
}


type GithubProfileResponse struct {
	Followers int `json:"followers"`
	Following int `json:"following"`
}

// function to get number of follower from github profile 
func (server *Server) GetNumberOfFollower (url string) (GithubProfileResponse, error) {

	// url := fmt.Sprintf("%s/users/%s", server.config.GithubEndpoint,  login)
	emptyResponse := GithubProfileResponse{
		Followers: 0,
		Following: 0,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return emptyResponse, err
	}

	req.Header =  http.Header{
    "Authorization": []string{ fmt.Sprintf("Bearer %s", server.config.GithubToken)},
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emptyResponse, err
	}

	// if resp status not 200 then retun error
	if resp.StatusCode != http.StatusOK	{
		// resp status not found then return eror
		if resp.StatusCode == http.StatusNotFound {
			return emptyResponse, fmt.Errorf("org not found")
		}
		return emptyResponse, fmt.Errorf("github api error")
	}

	defer resp.Body.Close()

	//Create a variable of the same type as our model
	 var cResp GithubProfileResponse

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return emptyResponse, err
	}

	return cResp, nil
}


type ListMemberOrgWithFollowers []struct {
   Login string    `json:"login"`
	 AvatarURL string `json:"avatar_url"`
	 Followers int `json:"followers"`
	 Following int `json:"following"`
}

// function to get number of follower from list of member with calling function GetNumberOfFollower
func (server *Server) GetNumberOfFollowerFromListMember (members ListMemberOrg) (ListMemberOrgWithFollowers, error) {
		membersWithFollowers := make(ListMemberOrgWithFollowers, len(members))
		for i, member := range members {
			followers, err := server.GetNumberOfFollower(member.URL)
			if err != nil {
				return nil, err
			}
			membersWithFollowers[i].Login = member.Login
			membersWithFollowers[i].AvatarURL = member.AvatarURL
			membersWithFollowers[i].Followers = followers.Followers
			membersWithFollowers[i].Following = followers.Following
		}

	return membersWithFollowers, nil
}
