// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/ent/chatmessage"
	"github.com/lxxonx/cinder-server/ent/chatroom"
	"github.com/lxxonx/cinder-server/ent/group"
	"github.com/lxxonx/cinder-server/ent/pic"
	"github.com/lxxonx/cinder-server/ent/schema"
	"github.com/lxxonx/cinder-server/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatmessageFields := schema.ChatMessage{}.Fields()
	_ = chatmessageFields
	// chatmessageDescMessage is the schema descriptor for message field.
	chatmessageDescMessage := chatmessageFields[1].Descriptor()
	// chatmessage.DefaultMessage holds the default value on creation for the message field.
	chatmessage.DefaultMessage = chatmessageDescMessage.Default.(string)
	// chatmessageDescCreatedAt is the schema descriptor for created_at field.
	chatmessageDescCreatedAt := chatmessageFields[4].Descriptor()
	// chatmessage.DefaultCreatedAt holds the default value on creation for the created_at field.
	chatmessage.DefaultCreatedAt = chatmessageDescCreatedAt.Default.(func() time.Time)
	// chatmessageDescUpdatedAt is the schema descriptor for updated_at field.
	chatmessageDescUpdatedAt := chatmessageFields[5].Descriptor()
	// chatmessage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	chatmessage.DefaultUpdatedAt = chatmessageDescUpdatedAt.Default.(func() time.Time)
	// chatmessage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	chatmessage.UpdateDefaultUpdatedAt = chatmessageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// chatmessageDescReadAt is the schema descriptor for read_at field.
	chatmessageDescReadAt := chatmessageFields[6].Descriptor()
	// chatmessage.DefaultReadAt holds the default value on creation for the read_at field.
	chatmessage.DefaultReadAt = chatmessageDescReadAt.Default.(func() time.Time)
	// chatmessageDescID is the schema descriptor for id field.
	chatmessageDescID := chatmessageFields[0].Descriptor()
	// chatmessage.DefaultID holds the default value on creation for the id field.
	chatmessage.DefaultID = chatmessageDescID.Default.(func() uuid.UUID)
	chatroomFields := schema.ChatRoom{}.Fields()
	_ = chatroomFields
	// chatroomDescCreatedAt is the schema descriptor for created_at field.
	chatroomDescCreatedAt := chatroomFields[1].Descriptor()
	// chatroom.DefaultCreatedAt holds the default value on creation for the created_at field.
	chatroom.DefaultCreatedAt = chatroomDescCreatedAt.Default.(func() time.Time)
	// chatroomDescUpdatedAt is the schema descriptor for updated_at field.
	chatroomDescUpdatedAt := chatroomFields[2].Descriptor()
	// chatroom.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	chatroom.DefaultUpdatedAt = chatroomDescUpdatedAt.Default.(func() time.Time)
	// chatroom.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	chatroom.UpdateDefaultUpdatedAt = chatroomDescUpdatedAt.UpdateDefault.(func() time.Time)
	// chatroomDescReadAt is the schema descriptor for read_at field.
	chatroomDescReadAt := chatroomFields[3].Descriptor()
	// chatroom.DefaultReadAt holds the default value on creation for the read_at field.
	chatroom.DefaultReadAt = chatroomDescReadAt.Default.(func() time.Time)
	// chatroomDescID is the schema descriptor for id field.
	chatroomDescID := chatroomFields[0].Descriptor()
	// chatroom.DefaultID holds the default value on creation for the id field.
	chatroom.DefaultID = chatroomDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescGroupname is the schema descriptor for groupname field.
	groupDescGroupname := groupFields[1].Descriptor()
	// group.GroupnameValidator is a validator for the "groupname" field. It is called by the builders before save.
	group.GroupnameValidator = groupDescGroupname.Validators[0].(func(string) error)
	// groupDescBio is the schema descriptor for bio field.
	groupDescBio := groupFields[2].Descriptor()
	// group.DefaultBio holds the default value on creation for the bio field.
	group.DefaultBio = groupDescBio.Default.(string)
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[3].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[4].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescReadAt is the schema descriptor for read_at field.
	groupDescReadAt := groupFields[5].Descriptor()
	// group.DefaultReadAt holds the default value on creation for the read_at field.
	group.DefaultReadAt = groupDescReadAt.Default.(func() time.Time)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	picFields := schema.Pic{}.Fields()
	_ = picFields
	// picDescURL is the schema descriptor for url field.
	picDescURL := picFields[3].Descriptor()
	// pic.URLValidator is a validator for the "url" field. It is called by the builders before save.
	pic.URLValidator = picDescURL.Validators[0].(func(string) error)
	// picDescCreatedAt is the schema descriptor for created_at field.
	picDescCreatedAt := picFields[4].Descriptor()
	// pic.DefaultCreatedAt holds the default value on creation for the created_at field.
	pic.DefaultCreatedAt = picDescCreatedAt.Default.(func() time.Time)
	// picDescUpdatedAt is the schema descriptor for updated_at field.
	picDescUpdatedAt := picFields[5].Descriptor()
	// pic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pic.DefaultUpdatedAt = picDescUpdatedAt.Default.(func() time.Time)
	// pic.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pic.UpdateDefaultUpdatedAt = picDescUpdatedAt.UpdateDefault.(func() time.Time)
	// picDescReadAt is the schema descriptor for read_at field.
	picDescReadAt := picFields[6].Descriptor()
	// pic.DefaultReadAt holds the default value on creation for the read_at field.
	pic.DefaultReadAt = picDescReadAt.Default.(func() time.Time)
	// picDescID is the schema descriptor for id field.
	picDescID := picFields[0].Descriptor()
	// pic.DefaultID holds the default value on creation for the id field.
	pic.DefaultID = picDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescActualName is the schema descriptor for actual_name field.
	userDescActualName := userFields[1].Descriptor()
	// user.ActualNameValidator is a validator for the "actual_name" field. It is called by the builders before save.
	user.ActualNameValidator = userDescActualName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[3].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescUni is the schema descriptor for uni field.
	userDescUni := userFields[5].Descriptor()
	// user.UniValidator is a validator for the "uni" field. It is called by the builders before save.
	user.UniValidator = userDescUni.Validators[0].(func(string) error)
	// userDescDep is the schema descriptor for dep field.
	userDescDep := userFields[6].Descriptor()
	// user.DepValidator is a validator for the "dep" field. It is called by the builders before save.
	user.DepValidator = userDescDep.Validators[0].(func(string) error)
	// userDescBirthYear is the schema descriptor for birth_year field.
	userDescBirthYear := userFields[8].Descriptor()
	// user.BirthYearValidator is a validator for the "birth_year" field. It is called by the builders before save.
	user.BirthYearValidator = func() func(int) error {
		validators := userDescBirthYear.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(birth_year int) error {
			for _, fn := range fns {
				if err := fn(birth_year); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescIsVerified is the schema descriptor for is_verified field.
	userDescIsVerified := userFields[9].Descriptor()
	// user.DefaultIsVerified holds the default value on creation for the is_verified field.
	user.DefaultIsVerified = userDescIsVerified.Default.(bool)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[10].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(string)
	// userDescMaxGroup is the schema descriptor for max_group field.
	userDescMaxGroup := userFields[11].Descriptor()
	// user.DefaultMaxGroup holds the default value on creation for the max_group field.
	user.DefaultMaxGroup = userDescMaxGroup.Default.(int)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[13].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[14].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescReadAt is the schema descriptor for read_at field.
	userDescReadAt := userFields[15].Descriptor()
	// user.DefaultReadAt holds the default value on creation for the read_at field.
	user.DefaultReadAt = userDescReadAt.Default.(func() time.Time)
}
