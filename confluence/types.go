package confluence
import 	"time"
type Content struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Space  struct {
		ID         int    `json:"id"`
		Key        string `json:"key"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Status     string `json:"status"`
		Expandable struct {
			Settings    string `json:"settings"`
			Metadata    string `json:"metadata"`
			Operations  string `json:"operations"`
			LookAndFeel string `json:"lookAndFeel"`
			Permissions string `json:"permissions"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
			Theme       string `json:"theme"`
			History     string `json:"history"`
			Homepage    string `json:"homepage"`
		} `json:"_expandable"`
		Links struct {
			Webui string `json:"webui"`
			Self  string `json:"self"`
		} `json:"_links"`
	} `json:"space"`
	History struct {
		Latest    bool `json:"latest"`
		CreatedBy struct {
			Type           string `json:"type"`
			Username       string `json:"username"`
			UserKey        string `json:"userKey"`
			AccountID      string `json:"accountId"`
			ProfilePicture struct {
				Path      string `json:"path"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsDefault bool   `json:"isDefault"`
			} `json:"profilePicture"`
			DisplayName string `json:"displayName"`
			Expandable  struct {
				Operations    string `json:"operations"`
				Details       string `json:"details"`
				PersonalSpace string `json:"personalSpace"`
			} `json:"_expandable"`
			Links struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"createdBy"`
		CreatedDate time.Time `json:"createdDate"`
		Expandable  struct {
			LastUpdated     string `json:"lastUpdated"`
			PreviousVersion string `json:"previousVersion"`
			Contributors    string `json:"contributors"`
			NextVersion     string `json:"nextVersion"`
		} `json:"_expandable"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"history"`
	Version struct {
		By struct {
			Type           string `json:"type"`
			Username       string `json:"username"`
			UserKey        string `json:"userKey"`
			AccountID      string `json:"accountId"`
			ProfilePicture struct {
				Path      string `json:"path"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsDefault bool   `json:"isDefault"`
			} `json:"profilePicture"`
			DisplayName string `json:"displayName"`
			Expandable  struct {
				Operations    string `json:"operations"`
				Details       string `json:"details"`
				PersonalSpace string `json:"personalSpace"`
			} `json:"_expandable"`
			Links struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"by"`
		When          time.Time `json:"when"`
		FriendlyWhen  string    `json:"friendlyWhen"`
		Message       string    `json:"message"`
		Number        int       `json:"number"`
		MinorEdit     bool      `json:"minorEdit"`
		SyncRev       string    `json:"syncRev"`
		SyncRevSource string    `json:"syncRevSource"`
		ConfRev       string    `json:"confRev"`
		Expandable    struct {
			Collaborators string `json:"collaborators"`
			Content       string `json:"content"`
		} `json:"_expandable"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"version"`
	MacroRenderedOutput struct {
	} `json:"macroRenderedOutput"`
	Extensions struct {
		Position int `json:"position"`
	} `json:"extensions"`
	Expandable struct {
		ChildTypes   string `json:"childTypes"`
		Container    string `json:"container"`
		Metadata     string `json:"metadata"`
		Operations   string `json:"operations"`
		Children     string `json:"children"`
		Restrictions string `json:"restrictions"`
		Ancestors    string `json:"ancestors"`
		Body         string `json:"body"`
		Descendants  string `json:"descendants"`
	} `json:"_expandable"`
	Links struct {
		Editui     string `json:"editui"`
		Webui      string `json:"webui"`
		Context    string `json:"context"`
		Self       string `json:"self"`
		Tinyui     string `json:"tinyui"`
		Collection string `json:"collection"`
		Base       string `json:"base"`
	} `json:"_links"`
}

type Version struct {
	Number int `json:"number"`
}

type Space struct {
	Key string `json:"key"`
}

type Body struct {
	View        *View `json:"view"`
	ExportView  *View `json:"export_view"`
	StyledView  *View `json:"styled_view"`
	StorageView *View `json:"storage"`
}

// View is data view
type View struct {
	Representation string `json:"representation"`
	Value          string `json:"value"`
}

type ContentRequest struct {
	Status  string   `json:"status"`
	Type    string   `json:"type"`
	Title   string   `json:"title"`
	Version *Version `json:"version"`
	Space   *Space   `json:"space"`
	Body    *Body    `json:"body"`
}
