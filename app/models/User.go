package models

import (
    "time"
)

// class
type User2 struct {
    ID              uint32      `gorm:"primary_key"`
    Name            string      `gorm:"size:20;not null"`
    Email           string      `gorm:"size:50;unique"`
    Password        string      `gorm:"size:255;not null"`
    InvitationCode  string      `gorm:"size:10"`
    QrCode          string      `gorm:"size:12"`
    RememberToken   string      `gorm:"size:100"`
    Comment         string      `gorm:"size:255;set utf8mb4"`
    TestUnique1     string      `gorm:"size:255;unique"`
    TestUnique2     string      `gorm:"size:255;unique_index"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       *time.Time
}
