package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileLobApp = MacOSLobApp{}

type MacOSLobApp struct {
	// The build number of the package. This should match the package CFBundleShortVersionString of the .pkg file.
	BuildNumber nullable.Type[string] `json:"buildNumber,omitempty"`

	// The primary bundleId of the package.
	BundleId nullable.Type[string] `json:"bundleId,omitempty"`

	// List of ComplexType macOSLobChildApp objects. Represents the apps expected to be installed by the package.
	ChildApps *[]MacOSLobChildApp `json:"childApps,omitempty"`

	// When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When
	// FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true
	// for apps that use a self update feature. The default value is FALSE.
	IgnoreVersionDetection *bool `json:"ignoreVersionDetection,omitempty"`

	// When TRUE, indicates that the app will be installed as managed (requires macOS 11.0 and other managed package
	// restrictions). When FALSE, indicates that the app will be installed as unmanaged. The default value is FALSE.
	InstallAsManaged *bool `json:"installAsManaged,omitempty"`

	// The MD5 hash codes. This is empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to
	// create a .intunemac, this value can be found inside the Detection.xml file.
	Md5Hash *[]string `json:"md5Hash,omitempty"`

	// The chunk size for MD5 hash. This is '0' or empty if the package was uploaded directly. If the Intune App Wrapping
	// Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
	Md5HashChunkSize *int64 `json:"md5HashChunkSize,omitempty"`

	// ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
	MinimumSupportedOperatingSystem *MacOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`

	// The version number of the package. This should match the package CFBundleVersion in the packageinfo file.
	VersionNumber nullable.Type[string] `json:"versionNumber,omitempty"`

	// Fields inherited from MobileLobApp

	// The internal committed content version.
	CommittedContentVersion nullable.Type[string] `json:"committedContentVersion,omitempty"`

	// The list of content versions for this app. This property is read-only.
	ContentVersions *[]MobileAppContent `json:"contentVersions,omitempty"`

	// The name of the main Lob application file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The total size, including all uploaded files. This property is read-only.
	Size *int64 `json:"size,omitempty"`

	// Fields inherited from MobileApp

	// The list of group assignments for this mobile app.
	Assignments *[]MobileAppAssignment `json:"assignments,omitempty"`

	// The list of categories for this app.
	Categories *[]MobileAppCategory `json:"categories,omitempty"`

	// The date and time the app was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The total number of dependencies the child app has.
	DependentAppCount *int64 `json:"dependentAppCount,omitempty"`

	// The description of the app.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The developer of the app.
	Developer nullable.Type[string] `json:"developer,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// The value indicating whether the app is assigned to at least one group.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Notes for the app.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The owner of the app.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The publisher of the app.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates the publishing state of an app.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`

	// List of relationships for this mobile app.
	Relationships *[]MobileAppRelationship `json:"relationships,omitempty"`

	// List of scope tag ids for this mobile app.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The total number of apps this app is directly or indirectly superseded by. This property is read-only.
	SupersededAppCount *int64 `json:"supersededAppCount,omitempty"`

	// The total number of apps this app directly or indirectly supersedes. This property is read-only.
	SupersedingAppCount *int64 `json:"supersedingAppCount,omitempty"`

	// The upload state.
	UploadState *int64 `json:"uploadState,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MacOSLobApp) MobileLobApp() BaseMobileLobAppImpl {
	return BaseMobileLobAppImpl{
		CommittedContentVersion: s.CommittedContentVersion,
		ContentVersions:         s.ContentVersions,
		FileName:                s.FileName,
		Size:                    s.Size,
		Assignments:             s.Assignments,
		Categories:              s.Categories,
		CreatedDateTime:         s.CreatedDateTime,
		DependentAppCount:       s.DependentAppCount,
		Description:             s.Description,
		Developer:               s.Developer,
		DisplayName:             s.DisplayName,
		InformationUrl:          s.InformationUrl,
		IsAssigned:              s.IsAssigned,
		IsFeatured:              s.IsFeatured,
		LargeIcon:               s.LargeIcon,
		LastModifiedDateTime:    s.LastModifiedDateTime,
		Notes:                   s.Notes,
		Owner:                   s.Owner,
		PrivacyInformationUrl:   s.PrivacyInformationUrl,
		Publisher:               s.Publisher,
		PublishingState:         s.PublishingState,
		Relationships:           s.Relationships,
		RoleScopeTagIds:         s.RoleScopeTagIds,
		SupersededAppCount:      s.SupersededAppCount,
		SupersedingAppCount:     s.SupersedingAppCount,
		UploadState:             s.UploadState,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s MacOSLobApp) MobileApp() BaseMobileAppImpl {
	return BaseMobileAppImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		DependentAppCount:     s.DependentAppCount,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsAssigned:            s.IsAssigned,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Relationships:         s.Relationships,
		RoleScopeTagIds:       s.RoleScopeTagIds,
		SupersededAppCount:    s.SupersededAppCount,
		SupersedingAppCount:   s.SupersedingAppCount,
		UploadState:           s.UploadState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s MacOSLobApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSLobApp{}

func (s MacOSLobApp) MarshalJSON() ([]byte, error) {
	type wrapper MacOSLobApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSLobApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSLobApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSLobApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSLobApp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MacOSLobApp{}

func (s *MacOSLobApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BuildNumber                     nullable.Type[string]        `json:"buildNumber,omitempty"`
		BundleId                        nullable.Type[string]        `json:"bundleId,omitempty"`
		ChildApps                       *[]MacOSLobChildApp          `json:"childApps,omitempty"`
		IgnoreVersionDetection          *bool                        `json:"ignoreVersionDetection,omitempty"`
		InstallAsManaged                *bool                        `json:"installAsManaged,omitempty"`
		Md5Hash                         *[]string                    `json:"md5Hash,omitempty"`
		Md5HashChunkSize                *int64                       `json:"md5HashChunkSize,omitempty"`
		MinimumSupportedOperatingSystem *MacOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
		VersionNumber                   nullable.Type[string]        `json:"versionNumber,omitempty"`
		CommittedContentVersion         nullable.Type[string]        `json:"committedContentVersion,omitempty"`
		ContentVersions                 *[]MobileAppContent          `json:"contentVersions,omitempty"`
		FileName                        nullable.Type[string]        `json:"fileName,omitempty"`
		Size                            *int64                       `json:"size,omitempty"`
		Assignments                     *[]MobileAppAssignment       `json:"assignments,omitempty"`
		Categories                      *[]MobileAppCategory         `json:"categories,omitempty"`
		CreatedDateTime                 *string                      `json:"createdDateTime,omitempty"`
		DependentAppCount               *int64                       `json:"dependentAppCount,omitempty"`
		Description                     nullable.Type[string]        `json:"description,omitempty"`
		Developer                       nullable.Type[string]        `json:"developer,omitempty"`
		DisplayName                     nullable.Type[string]        `json:"displayName,omitempty"`
		InformationUrl                  nullable.Type[string]        `json:"informationUrl,omitempty"`
		IsAssigned                      *bool                        `json:"isAssigned,omitempty"`
		IsFeatured                      *bool                        `json:"isFeatured,omitempty"`
		LargeIcon                       *MimeContent                 `json:"largeIcon,omitempty"`
		LastModifiedDateTime            *string                      `json:"lastModifiedDateTime,omitempty"`
		Notes                           nullable.Type[string]        `json:"notes,omitempty"`
		Owner                           nullable.Type[string]        `json:"owner,omitempty"`
		PrivacyInformationUrl           nullable.Type[string]        `json:"privacyInformationUrl,omitempty"`
		Publisher                       nullable.Type[string]        `json:"publisher,omitempty"`
		PublishingState                 *MobileAppPublishingState    `json:"publishingState,omitempty"`
		RoleScopeTagIds                 *[]string                    `json:"roleScopeTagIds,omitempty"`
		SupersededAppCount              *int64                       `json:"supersededAppCount,omitempty"`
		SupersedingAppCount             *int64                       `json:"supersedingAppCount,omitempty"`
		UploadState                     *int64                       `json:"uploadState,omitempty"`
		Id                              *string                      `json:"id,omitempty"`
		ODataId                         *string                      `json:"@odata.id,omitempty"`
		ODataType                       *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BuildNumber = decoded.BuildNumber
	s.BundleId = decoded.BundleId
	s.ChildApps = decoded.ChildApps
	s.IgnoreVersionDetection = decoded.IgnoreVersionDetection
	s.InstallAsManaged = decoded.InstallAsManaged
	s.Md5Hash = decoded.Md5Hash
	s.Md5HashChunkSize = decoded.Md5HashChunkSize
	s.MinimumSupportedOperatingSystem = decoded.MinimumSupportedOperatingSystem
	s.VersionNumber = decoded.VersionNumber
	s.Assignments = decoded.Assignments
	s.Categories = decoded.Categories
	s.CommittedContentVersion = decoded.CommittedContentVersion
	s.ContentVersions = decoded.ContentVersions
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DependentAppCount = decoded.DependentAppCount
	s.Description = decoded.Description
	s.Developer = decoded.Developer
	s.DisplayName = decoded.DisplayName
	s.FileName = decoded.FileName
	s.Id = decoded.Id
	s.InformationUrl = decoded.InformationUrl
	s.IsAssigned = decoded.IsAssigned
	s.IsFeatured = decoded.IsFeatured
	s.LargeIcon = decoded.LargeIcon
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Notes = decoded.Notes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Owner = decoded.Owner
	s.PrivacyInformationUrl = decoded.PrivacyInformationUrl
	s.Publisher = decoded.Publisher
	s.PublishingState = decoded.PublishingState
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Size = decoded.Size
	s.SupersededAppCount = decoded.SupersededAppCount
	s.SupersedingAppCount = decoded.SupersedingAppCount
	s.UploadState = decoded.UploadState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MacOSLobApp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["relationships"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Relationships into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileAppRelationship, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileAppRelationshipImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Relationships' for 'MacOSLobApp': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Relationships = &output
	}

	return nil
}
