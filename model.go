package orgperm

import "github.com/google/uuid"

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

type UpdatePermission struct {
	OrgUpdate     bool `json:"org_update"`
	RolesManage   bool `json:"roles_manage"`
	InvitesManage bool `json:"invites_manage"`

	MembersDelete bool `json:"members_delete"`
	MembersUpdate bool `json:"members_update"`

	PlacesCreate bool `json:"places_create"`
	PlacesDelete bool `json:"places_delete"`
	PlacesUpdate bool `json:"places_update"`
}

func (p UpdatePermission) ToCodeMap() map[string]bool {
	perms := make(map[string]bool)

	perms[OrganizationUpdateCode] = p.OrgUpdate
	perms[RolesManageCode] = p.RolesManage
	perms[InvitesManageCode] = p.InvitesManage

	perms[MembersDeleteCode] = p.MembersDelete
	perms[MembersUpdateCode] = p.MembersUpdate

	perms[PlacesCreateCode] = p.PlacesCreate
	perms[PlacesDeleteCode] = p.PlacesDelete
	perms[PlacesUpdateCode] = p.PlacesUpdate

	return perms
}

func (p UpdatePermission) ToIDMap() map[uuid.UUID]bool {
	perms := make(map[uuid.UUID]bool)

	perms[OrganizationUpdateID] = p.OrgUpdate
	perms[RolesManageID] = p.RolesManage
	perms[InvitesManageID] = p.InvitesManage

	perms[MembersDeleteID] = p.MembersDelete
	perms[MembersUpdateID] = p.MembersUpdate

	perms[PlacesCreateID] = p.PlacesCreate
	perms[PlacesDeleteID] = p.PlacesDelete
	perms[PlacesUpdateID] = p.PlacesUpdate

	return perms
}
