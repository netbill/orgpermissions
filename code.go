package orgperm

import "github.com/google/uuid"

var (
	OrganizationUpdateID = uuid.MustParse("8a1f6a63-2a5d-4b35-9f64-2f2c3bce9e01")
	RolesManageID        = uuid.MustParse("7d2c4e91-1c5a-4b23-a9b1-93df0a12f402")
	InvitesManageID      = uuid.MustParse("c1e0a572-9a2f-4e9c-b67e-5d01a7c44d03")

	MembersDeleteID = uuid.MustParse("2b1d5c64-6f12-4c78-9c34-8b2e6d0f1a04")
	MembersUpdateID = uuid.MustParse("a4e7c903-3d11-4b8f-8e1d-12f6b2d9c205")

	PlacesCreateID = uuid.MustParse("d6f3a890-7b2c-4e21-b93a-56a1c8d0e306")
	PlacesDeleteID = uuid.MustParse("e7a4b901-8c3d-4f12-a92b-67b2d9e1f407")
	PlacesUpdateID = uuid.MustParse("f8b5c012-9d4e-4a23-b81c-78c3e0f2a508")
)

const (
	OrganizationUpdateCode = "organization.update"
	RolesManageCode        = "roles.manage"
	InvitesManageCode      = "invites.manage"

	MembersDeleteCode = "members.delete"
	MembersUpdateCode = "members.update"

	PlacesCreateCode = "places.create"
	PlacesDeleteCode = "places.delete"
	PlacesUpdateCode = "places.update"
)

type Permission struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
}

func GetAllPermissions() []Permission {
	return []Permission{
		{ID: OrganizationUpdateID, Code: OrganizationUpdateCode},
		{ID: RolesManageID, Code: RolesManageCode},
		{ID: InvitesManageID, Code: InvitesManageCode},

		{ID: MembersDeleteID, Code: MembersDeleteCode},
		{ID: MembersUpdateID, Code: MembersUpdateCode},

		{ID: PlacesCreateID, Code: PlacesCreateCode},
		{ID: PlacesDeleteID, Code: PlacesDeleteCode},
		{ID: PlacesUpdateID, Code: PlacesUpdateCode},
	}
}

var NumOfAllPermissions = len(GetAllPermissions())
