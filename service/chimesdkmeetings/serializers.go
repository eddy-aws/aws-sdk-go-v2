// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpBatchCreateAttendee struct {
}

func (*awsRestjson1_serializeOpBatchCreateAttendee) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBatchCreateAttendee) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchCreateAttendeeInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/attendees?operation=batch-create")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsBatchCreateAttendeeInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentBatchCreateAttendeeInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsBatchCreateAttendeeInput(v *BatchCreateAttendeeInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBatchCreateAttendeeInput(v *BatchCreateAttendeeInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Attendees != nil {
		ok := object.Key("Attendees")
		if err := awsRestjson1_serializeDocumentCreateAttendeeRequestItemList(v.Attendees, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpCreateAttendee struct {
}

func (*awsRestjson1_serializeOpCreateAttendee) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateAttendee) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateAttendeeInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/attendees")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsCreateAttendeeInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateAttendeeInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateAttendeeInput(v *CreateAttendeeInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateAttendeeInput(v *CreateAttendeeInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExternalUserId != nil {
		ok := object.Key("ExternalUserId")
		ok.String(*v.ExternalUserId)
	}

	return nil
}

type awsRestjson1_serializeOpCreateMeeting struct {
}

func (*awsRestjson1_serializeOpCreateMeeting) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateMeeting) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateMeetingInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateMeetingInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateMeetingInput(v *CreateMeetingInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateMeetingInput(v *CreateMeetingInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientRequestToken != nil {
		ok := object.Key("ClientRequestToken")
		ok.String(*v.ClientRequestToken)
	}

	if v.ExternalMeetingId != nil {
		ok := object.Key("ExternalMeetingId")
		ok.String(*v.ExternalMeetingId)
	}

	if v.MediaRegion != nil {
		ok := object.Key("MediaRegion")
		ok.String(*v.MediaRegion)
	}

	if v.MeetingHostId != nil {
		ok := object.Key("MeetingHostId")
		ok.String(*v.MeetingHostId)
	}

	if v.NotificationsConfiguration != nil {
		ok := object.Key("NotificationsConfiguration")
		if err := awsRestjson1_serializeDocumentNotificationsConfiguration(v.NotificationsConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpCreateMeetingWithAttendees struct {
}

func (*awsRestjson1_serializeOpCreateMeetingWithAttendees) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateMeetingWithAttendees) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateMeetingWithAttendeesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings?operation=create-attendees")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateMeetingWithAttendeesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateMeetingWithAttendeesInput(v *CreateMeetingWithAttendeesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateMeetingWithAttendeesInput(v *CreateMeetingWithAttendeesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Attendees != nil {
		ok := object.Key("Attendees")
		if err := awsRestjson1_serializeDocumentCreateMeetingWithAttendeesRequestItemList(v.Attendees, ok); err != nil {
			return err
		}
	}

	if v.ClientRequestToken != nil {
		ok := object.Key("ClientRequestToken")
		ok.String(*v.ClientRequestToken)
	}

	if v.ExternalMeetingId != nil {
		ok := object.Key("ExternalMeetingId")
		ok.String(*v.ExternalMeetingId)
	}

	if v.MediaRegion != nil {
		ok := object.Key("MediaRegion")
		ok.String(*v.MediaRegion)
	}

	if v.MeetingHostId != nil {
		ok := object.Key("MeetingHostId")
		ok.String(*v.MeetingHostId)
	}

	if v.NotificationsConfiguration != nil {
		ok := object.Key("NotificationsConfiguration")
		if err := awsRestjson1_serializeDocumentNotificationsConfiguration(v.NotificationsConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDeleteAttendee struct {
}

func (*awsRestjson1_serializeOpDeleteAttendee) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteAttendee) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteAttendeeInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/attendees/{AttendeeId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteAttendeeInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteAttendeeInput(v *DeleteAttendeeInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AttendeeId == nil || len(*v.AttendeeId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member AttendeeId must not be empty")}
	}
	if v.AttendeeId != nil {
		if err := encoder.SetURI("AttendeeId").String(*v.AttendeeId); err != nil {
			return err
		}
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDeleteMeeting struct {
}

func (*awsRestjson1_serializeOpDeleteMeeting) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteMeeting) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteMeetingInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteMeetingInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteMeetingInput(v *DeleteMeetingInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetAttendee struct {
}

func (*awsRestjson1_serializeOpGetAttendee) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetAttendee) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetAttendeeInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/attendees/{AttendeeId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetAttendeeInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetAttendeeInput(v *GetAttendeeInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AttendeeId == nil || len(*v.AttendeeId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member AttendeeId must not be empty")}
	}
	if v.AttendeeId != nil {
		if err := encoder.SetURI("AttendeeId").String(*v.AttendeeId); err != nil {
			return err
		}
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetMeeting struct {
}

func (*awsRestjson1_serializeOpGetMeeting) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetMeeting) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetMeetingInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetMeetingInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetMeetingInput(v *GetMeetingInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListAttendees struct {
}

func (*awsRestjson1_serializeOpListAttendees) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListAttendees) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListAttendeesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/attendees")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListAttendeesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListAttendeesInput(v *ListAttendeesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("max-results").Integer(*v.MaxResults)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	if v.NextToken != nil {
		encoder.SetQuery("next-token").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpStartMeetingTranscription struct {
}

func (*awsRestjson1_serializeOpStartMeetingTranscription) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartMeetingTranscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartMeetingTranscriptionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/transcription?operation=start")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStartMeetingTranscriptionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartMeetingTranscriptionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartMeetingTranscriptionInput(v *StartMeetingTranscriptionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartMeetingTranscriptionInput(v *StartMeetingTranscriptionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TranscriptionConfiguration != nil {
		ok := object.Key("TranscriptionConfiguration")
		if err := awsRestjson1_serializeDocumentTranscriptionConfiguration(v.TranscriptionConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStopMeetingTranscription struct {
}

func (*awsRestjson1_serializeOpStopMeetingTranscription) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStopMeetingTranscription) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StopMeetingTranscriptionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/meetings/{MeetingId}/transcription?operation=stop")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStopMeetingTranscriptionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStopMeetingTranscriptionInput(v *StopMeetingTranscriptionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MeetingId == nil || len(*v.MeetingId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member MeetingId must not be empty")}
	}
	if v.MeetingId != nil {
		if err := encoder.SetURI("MeetingId").String(*v.MeetingId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentCreateAttendeeRequestItem(v *types.CreateAttendeeRequestItem, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExternalUserId != nil {
		ok := object.Key("ExternalUserId")
		ok.String(*v.ExternalUserId)
	}

	return nil
}

func awsRestjson1_serializeDocumentCreateAttendeeRequestItemList(v []types.CreateAttendeeRequestItem, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentCreateAttendeeRequestItem(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentCreateMeetingWithAttendeesRequestItemList(v []types.CreateAttendeeRequestItem, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentCreateAttendeeRequestItem(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentEngineTranscribeMedicalSettings(v *types.EngineTranscribeMedicalSettings, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.ContentIdentificationType) > 0 {
		ok := object.Key("ContentIdentificationType")
		ok.String(string(v.ContentIdentificationType))
	}

	if len(v.LanguageCode) > 0 {
		ok := object.Key("LanguageCode")
		ok.String(string(v.LanguageCode))
	}

	if len(v.Region) > 0 {
		ok := object.Key("Region")
		ok.String(string(v.Region))
	}

	if len(v.Specialty) > 0 {
		ok := object.Key("Specialty")
		ok.String(string(v.Specialty))
	}

	if len(v.Type) > 0 {
		ok := object.Key("Type")
		ok.String(string(v.Type))
	}

	if v.VocabularyName != nil {
		ok := object.Key("VocabularyName")
		ok.String(*v.VocabularyName)
	}

	return nil
}

func awsRestjson1_serializeDocumentEngineTranscribeSettings(v *types.EngineTranscribeSettings, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.ContentIdentificationType) > 0 {
		ok := object.Key("ContentIdentificationType")
		ok.String(string(v.ContentIdentificationType))
	}

	if len(v.ContentRedactionType) > 0 {
		ok := object.Key("ContentRedactionType")
		ok.String(string(v.ContentRedactionType))
	}

	if v.EnablePartialResultsStabilization {
		ok := object.Key("EnablePartialResultsStabilization")
		ok.Boolean(v.EnablePartialResultsStabilization)
	}

	if len(v.LanguageCode) > 0 {
		ok := object.Key("LanguageCode")
		ok.String(string(v.LanguageCode))
	}

	if v.LanguageModelName != nil {
		ok := object.Key("LanguageModelName")
		ok.String(*v.LanguageModelName)
	}

	if len(v.PartialResultsStability) > 0 {
		ok := object.Key("PartialResultsStability")
		ok.String(string(v.PartialResultsStability))
	}

	if v.PiiEntityTypes != nil {
		ok := object.Key("PiiEntityTypes")
		ok.String(*v.PiiEntityTypes)
	}

	if len(v.Region) > 0 {
		ok := object.Key("Region")
		ok.String(string(v.Region))
	}

	if len(v.VocabularyFilterMethod) > 0 {
		ok := object.Key("VocabularyFilterMethod")
		ok.String(string(v.VocabularyFilterMethod))
	}

	if v.VocabularyFilterName != nil {
		ok := object.Key("VocabularyFilterName")
		ok.String(*v.VocabularyFilterName)
	}

	if v.VocabularyName != nil {
		ok := object.Key("VocabularyName")
		ok.String(*v.VocabularyName)
	}

	return nil
}

func awsRestjson1_serializeDocumentNotificationsConfiguration(v *types.NotificationsConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.LambdaFunctionArn != nil {
		ok := object.Key("LambdaFunctionArn")
		ok.String(*v.LambdaFunctionArn)
	}

	if v.SnsTopicArn != nil {
		ok := object.Key("SnsTopicArn")
		ok.String(*v.SnsTopicArn)
	}

	if v.SqsQueueArn != nil {
		ok := object.Key("SqsQueueArn")
		ok.String(*v.SqsQueueArn)
	}

	return nil
}

func awsRestjson1_serializeDocumentTranscriptionConfiguration(v *types.TranscriptionConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EngineTranscribeMedicalSettings != nil {
		ok := object.Key("EngineTranscribeMedicalSettings")
		if err := awsRestjson1_serializeDocumentEngineTranscribeMedicalSettings(v.EngineTranscribeMedicalSettings, ok); err != nil {
			return err
		}
	}

	if v.EngineTranscribeSettings != nil {
		ok := object.Key("EngineTranscribeSettings")
		if err := awsRestjson1_serializeDocumentEngineTranscribeSettings(v.EngineTranscribeSettings, ok); err != nil {
			return err
		}
	}

	return nil
}
