package server

import (
	"context"
	"errors"
	"sync"

	"note-server/pb"
	"note-server/src/db"
	"note-server/src/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type NoteServer struct {
	pb.UnimplementedNoteServiceServer

	/**
	QUESTION
	How does this mutex prevent everything from being written to or read from?
	Everywhere in the code it just locks the mutex and defers the unlock and
	that's supposed to lock the whole struct?
	*/
	mutex sync.RWMutex
}

func NewNoteServer() *NoteServer {
	return &NoteServer{}
}

func (s *NoteServer) getNotesFromTitles(linked_note_titles []string) []*pb.Note {
	// linked_notes := []*pb.Note{}

	// for _, note_title := range linked_note_titles {
	// 	linked_note, exists := s.notes[note_title]
	// 	if exists {
	// 		linked_notes = append(linked_notes, linked_note)
	// 	}
	// }

	// return linked_notes
	return []*pb.Note{}
}

func (s *NoteServer) CreateNote(_ context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	if req.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "title cannot be empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	var count int64
	if err := db.DB.Model(&models.NoteModel{}).Where("title = ?", req.Title).Count(&count).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "an unknown error occurred during note counting: %w", err)
	}

	if count > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "note with title '%s' already exists", req.Title)
	}

	linked_notes := s.getNotesFromTitles(req.LinkedNoteTitles)
	if len(req.LinkedNoteTitles) != len(linked_notes) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"tried to find %d notes to link but only found %d notes in system",
			len(req.LinkedNoteTitles),
			len(linked_notes),
		)
	}

	newNote := &models.NoteModel{
		Title:   req.Title,
		Content: req.Content,
	}
	if err := db.DB.Create(&newNote).Error; err != nil {
		return nil, status.Errorf(codes.Unknown, "an unknown error occurred during note creation: %w", err)
	}

	response := &pb.CreateNoteResponse{
		Note: &pb.Note{
			Title:     newNote.Title,
			Content:   newNote.Content,
			CreatedAt: newNote.CreatedAt.Unix(),
			UpdatedAt: newNote.UpdatedAt.Unix(),
		},
	}

	return response, nil
}

func (s *NoteServer) GetNote(_ context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	if req.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "note Title cannot be empty")
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var noteModel models.NoteModel
	result := db.DB.First(&noteModel, "title = ?", req.Title)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, status.Errorf(codes.Unknown, "an unknown error occurred: %w", result.Error)
		}
	}

	note := &pb.Note{
		Title:     noteModel.Title,
		Content:   noteModel.Content,
		CreatedAt: noteModel.CreatedAt.Unix(),
		UpdatedAt: noteModel.UpdatedAt.Unix(),
	}

	response := &pb.GetNoteResponse{Note: note}

	return response, nil
}

func (s *NoteServer) ListNotes(_ context.Context, __ *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	// s.mutex.RLock()
	// defer s.mutex.RUnlock()

	return nil, status.Errorf(codes.Unimplemented, "this function is not implemented yet")
}

func (s *NoteServer) UpdateNote(_ context.Context, req *pb.UpdateNoteRequest) (*pb.UpdateNoteResponse, error) {
	// if req.Title == "" {
	// 	return nil, status.Errorf(codes.InvalidArgument, "title cannot be empty")
	// }
	// if req.Content == "" {
	// 	return nil, status.Errorf(codes.InvalidArgument, "content cannot be empty")
	// }

	// s.mutex.Lock()
	// defer s.mutex.Unlock()

	// note, exists := s.notes[req.Title]
	// if !exists {
	// 	return nil, status.Errorf(codes.NotFound, "note with ID %s not found", req.Title)
	// }

	// linked_notes := s.getNotesFromTitles(req.LinkedNoteTitles)
	// if len(req.LinkedNoteTitles) != len(linked_notes) {
	// 	return nil, status.Errorf(
	// 		codes.InvalidArgument,
	// 		"tried to link %d notes but only found %d notes in system",
	// 		len(req.LinkedNoteTitles),
	// 		len(linked_notes),
	// 	)
	// }

	// note.Title = req.Title
	// note.Content = req.Content
	// note.LinkedNotes = linked_notes
	// note.UpdatedAt = time.Now().Unix()
	// if req.Tags != nil {
	// 	note.Tags = req.Tags
	// }

	// response := &pb.UpdateNoteResponse{Note: note}

	// return response, nil

	return nil, status.Errorf(codes.Unimplemented, "this function is not implemented yet")
}

func (s *NoteServer) DeleteNote(_ context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	// if req.Title == "" {
	// 	return nil, status.Errorf(codes.InvalidArgument, "note Title cannot be empty")
	// }

	// s.mutex.Lock()
	// defer s.mutex.Unlock()

	// // TODO: should this really return an error if it's already been deleted?
	// _, exists := s.notes[req.Title]
	// if !exists {
	// 	return nil, status.Errorf(codes.NotFound, "note with ID %s not found", req.Title)
	// }

	// delete(s.notes, req.Title)

	// response := &pb.DeleteNoteResponse{Success: true}

	// return response, nil

	return nil, status.Errorf(codes.Unimplemented, "this function is not implemented yet")
}
