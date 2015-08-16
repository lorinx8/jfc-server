package nfutil

import (
	"testing"
)

func TestPutLocalToCloud(t *testing.T) {
	data := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01, 0x01, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0xFF, 0xDB, 0x00, 0x43, 0x00, 0x02, 0x01, 0x01, 0x01, 0x01, 0x01, 0x02, 0x01, 0x01, 0x01, 0x02, 0x02, 0x02, 0x02, 0x02, 0x04, 0x03, 0x02, 0x02, 0x02, 0x02, 0x05, 0x04, 0x04, 0x03, 0x04, 0x06, 0x05, 0x06, 0x06, 0x06, 0x05, 0x06, 0x06, 0x06, 0x07, 0x09, 0x08, 0x06, 0x07, 0x09, 0x07, 0x06, 0x06, 0x08, 0x0B, 0x08, 0x09, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x06, 0x08, 0x0B, 0x0C, 0x0B, 0x0A, 0x0C, 0x09, 0x0A, 0x0A, 0x0A, 0xFF, 0xDB, 0x00, 0x43, 0x01, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x05, 0x03, 0x03, 0x05, 0x0A, 0x07, 0x06, 0x07, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0xFF, 0xC0, 0x00, 0x11, 0x08, 0x00, 0x24, 0x00, 0x88, 0x03, 0x01, 0x22, 0x00, 0x02, 0x11, 0x01, 0x03, 0x11, 0x01, 0xFF, 0xC4, 0x00, 0x1F, 0x00, 0x00, 0x01, 0x05, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0xFF, 0xC4, 0x00, 0xB5, 0x10, 0x00, 0x02, 0x01, 0x03, 0x03, 0x02, 0x04, 0x03, 0x05, 0x05, 0x04, 0x04, 0x00, 0x00, 0x01, 0x7D, 0x01, 0x02, 0x03, 0x00, 0x04, 0x11, 0x05, 0x12, 0x21, 0x31, 0x41, 0x06, 0x13, 0x51, 0x61, 0x07, 0x22, 0x71, 0x14, 0x32, 0x81, 0x91, 0xA1, 0x08, 0x23, 0x42, 0xB1, 0xC1, 0x15, 0x52, 0xD1, 0xF0, 0x24, 0x33, 0x62, 0x72, 0x82, 0x09, 0x0A, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFF, 0xC4, 0x00, 0x1F, 0x01, 0x00, 0x03, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0xFF, 0xC4, 0x00, 0xB5, 0x11, 0x00, 0x02, 0x01, 0x02, 0x04, 0x04, 0x03, 0x04, 0x07, 0x05, 0x04, 0x04, 0x00, 0x01, 0x02, 0x77, 0x00, 0x01, 0x02, 0x03, 0x11, 0x04, 0x05, 0x21, 0x31, 0x06, 0x12, 0x41, 0x51, 0x07, 0x61, 0x71, 0x13, 0x22, 0x32, 0x81, 0x08, 0x14, 0x42, 0x91, 0xA1, 0xB1, 0xC1, 0x09, 0x23, 0x33, 0x52, 0xF0, 0x15, 0x62, 0x72, 0xD1, 0x0A, 0x16, 0x24, 0x34, 0xE1, 0x25, 0xF1, 0x17, 0x18, 0x19, 0x1A, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFF, 0xDA, 0x00, 0x0C, 0x03, 0x01, 0x00, 0x02, 0x11, 0x03, 0x11, 0x00, 0x3F, 0x00, 0xE3, 0x75, 0x8F, 0xDA, 0x93, 0xC5, 0xBF, 0xDA, 0x13, 0x7F, 0x62, 0x58, 0x5B, 0x43, 0x6E, 0x24, 0x61, 0x08, 0x78, 0xF2, 0xD8, 0x07, 0x1C, 0xD5, 0x76, 0xFD, 0xA6, 0xFE, 0x2A, 0x46, 0x00, 0x92, 0x78, 0xA0, 0xDE, 0xBB, 0x90, 0xB5, 0xB6, 0x37, 0x0F, 0x6C, 0x8E, 0x6A, 0x7D, 0x39, 0xF4, 0x8F, 0x88, 0x1F, 0x0A, 0xF4, 0x3F, 0x83, 0xE9, 0xE0, 0x43, 0x6B, 0xE2, 0x1B, 0xED, 0x54, 0x8D, 0x13, 0x5C, 0x68, 0x4A, 0x47, 0x3A, 0x16, 0x39, 0xCB, 0xE3, 0xE6, 0xAF, 0x5B, 0xF8, 0xA5, 0xFB, 0x24, 0x7C, 0x45, 0xF1, 0xB7, 0x82, 0xF4, 0xAF, 0x0F, 0x4D, 0xE2, 0x4F, 0x0A, 0x5B, 0xEA, 0x9E, 0x14, 0xD3, 0xA4, 0xF3, 0xD2, 0xD1, 0x8B, 0x5C, 0x5C, 0x22, 0xAE, 0xEC, 0x32, 0x83, 0xC7, 0x4E, 0xB5, 0xFB, 0x16, 0x0F, 0x2A, 0xC9, 0x69, 0xC2, 0x3C, 0xD4, 0xB5, 0x67, 0xCA, 0x54, 0xC7, 0xD7, 0x8E, 0x8D, 0xA3, 0xC5, 0xAE, 0xBF, 0x68, 0xEF, 0x8A, 0xAC, 0x82, 0x5F, 0xED, 0xA4, 0x50, 0x3D, 0x20, 0x15, 0xE8, 0xDE, 0x08, 0xF8, 0xFD, 0xE3, 0xAF, 0x00, 0x69, 0xB6, 0xFA, 0xFF, 0x00, 0xC4, 0xAD, 0x16, 0xF6, 0xD2, 0x1D, 0x4E, 0x26, 0x4B, 0x2B, 0xF9, 0xA0, 0x1E, 0x54, 0xFC, 0x67, 0x8F, 0x7E, 0x95, 0xF3, 0xDE, 0x99, 0x65, 0x75, 0xAC, 0xEB, 0xD6, 0x3A, 0x28, 0x04, 0xCB, 0x3E, 0xA1, 0x1C, 0x0C, 0x17, 0xF8, 0xBE, 0x7C, 0x1E, 0x2B, 0xEE, 0x6F, 0xDB, 0xAF, 0xE1, 0x75, 0xB0, 0xFD, 0x90, 0xF4, 0xE6, 0xB4, 0xD3, 0x82, 0xCD, 0xE1, 0xE1, 0x04, 0x91, 0x6C, 0x5E, 0x71, 0x80, 0x0D, 0x74, 0x66, 0x99, 0x4E, 0x55, 0x07, 0x0A, 0x50, 0x85, 0x9C, 0x8E, 0x59, 0x66, 0x72, 0x6B, 0xDE, 0x3E, 0x7E, 0xF1, 0x07, 0xED, 0x33, 0x14, 0x7A, 0x95, 0xD5, 0xFF, 0x00, 0x87, 0xED, 0x9E, 0xED, 0xAE, 0xA5, 0xDE, 0xCD, 0x7C, 0x72, 0xA9, 0xC7, 0x21, 0x40, 0xE8, 0x2B, 0x2E, 0x4F, 0xDA, 0xAB, 0xC6, 0x71, 0xC8, 0xAD, 0x69, 0x63, 0x63, 0x1B, 0x74, 0x1B, 0x60, 0x27, 0x22, 0xBC, 0xAE, 0xDA, 0x35, 0x16, 0xE8, 0x49, 0xE4, 0x81, 0xCF, 0xAD, 0x77, 0x9F, 0xB3, 0x07, 0x82, 0xB4, 0xBF, 0x88, 0x5F, 0xB4, 0x27, 0x87, 0x3C, 0x29, 0xAE, 0x7C, 0xD6, 0x72, 0xDC, 0x99, 0x25, 0x56, 0x3C, 0x36, 0xD1, 0x90, 0xBF, 0x8D, 0x77, 0x43, 0x85, 0x72, 0x6C, 0x26, 0x1D, 0xCE, 0x70, 0xBB, 0x5A, 0x99, 0xBC, 0x65, 0xF6, 0x47, 0xAA, 0x78, 0x76, 0xFF, 0x00, 0xF6, 0xC4, 0xF1, 0x96, 0x80, 0x7C, 0x4B, 0xE1, 0xAF, 0x01, 0x16, 0xB3, 0x0B, 0xB9, 0x5C, 0xC5, 0xB4, 0xB0, 0xF5, 0x00, 0xD7, 0x05, 0xAA, 0xFE, 0xD1, 0x9F, 0x1A, 0x46, 0xA8, 0x7C, 0x39, 0xAA, 0x5E, 0xB5, 0xA5, 0xEA, 0xCE, 0x21, 0x36, 0xF2, 0xC3, 0xB4, 0xAB, 0x13, 0x81, 0xFA, 0xD7, 0xA8, 0xFC, 0x7E, 0xFD, 0xAD, 0xFE, 0x22, 0xF8, 0x03, 0xF6, 0x84, 0xFF, 0x00, 0x84, 0x2B, 0xC1, 0xCC, 0x6C, 0xB4, 0xFD, 0x1E, 0xEE, 0x3B, 0x78, 0xF4, 0xE4, 0x5C, 0x09, 0x06, 0x40, 0xE9, 0xDE, 0xB2, 0x7F, 0x6F, 0x0D, 0x0F, 0x4C, 0x1F, 0x19, 0xBC, 0x1B, 0xE2, 0xAB, 0x5B, 0x48, 0xA1, 0xB9, 0xD6, 0x22, 0xB7, 0x92, 0xF6, 0x18, 0xD7, 0x18, 0x72, 0x54, 0xF2, 0x2B, 0xCE, 0xC3, 0xE0, 0x70, 0x10, 0x9F, 0xBD, 0x45, 0x5B, 0x5E, 0xC2, 0x8E, 0x2E, 0xB4, 0xB6, 0x39, 0x8F, 0x8A, 0x9E, 0x2F, 0xFD, 0xA2, 0x3E, 0x0D, 0xEA, 0x56, 0x7A, 0x5F, 0x8C, 0x75, 0xF5, 0x59, 0x2F, 0xED, 0x56, 0x7B, 0x7F, 0x23, 0xB2, 0x91, 0x9E, 0x6B, 0x91, 0xB8, 0xFD, 0xA0, 0xBE, 0x2B, 0xDF, 0x00, 0xAF, 0xE2, 0xFB, 0xA4, 0x1E, 0x8A, 0xF8, 0xAF, 0x56, 0xFF, 0x00, 0x82, 0x8F, 0x30, 0x5F, 0x1E, 0xF8, 0x66, 0x04, 0xCE, 0x53, 0x41, 0x8B, 0x24, 0xFF, 0x00, 0xBA, 0x2B, 0xE7, 0x45, 0xE4, 0xE3, 0x77, 0x3E, 0xA4, 0x57, 0xB3, 0x95, 0xE5, 0x39, 0x6D, 0x6C, 0x3F, 0x33, 0xA4, 0x82, 0x78, 0xC9, 0xC5, 0x6A, 0xCE, 0xB2, 0x7F, 0x8D, 0xDF, 0x13, 0x16, 0x2C, 0x3F, 0x8D, 0x2F, 0x0A, 0x9F, 0xBC, 0x44, 0xA7, 0x8A, 0xF4, 0x6F, 0x05, 0x7C, 0x30, 0xF8, 0xB9, 0xE3, 0x5F, 0x80, 0xFA, 0xCF, 0xC7, 0x8D, 0x47, 0xE2, 0x26, 0xA1, 0x67, 0x05, 0x8E, 0xEF, 0xB2, 0x5A, 0x99, 0x98, 0xF9, 0xEA, 0x3F, 0x8A, 0xBC, 0x87, 0xC0, 0xBE, 0x0D, 0xD5, 0xFE, 0x22, 0xF8, 0xDB, 0x4C, 0xF0, 0x2E, 0x91, 0x13, 0x3C, 0xDA, 0x95, 0xD2, 0xC6, 0xCC, 0x8B, 0x9D, 0xA9, 0x9F, 0x98, 0xFB, 0x57, 0xDD, 0x9F, 0x19, 0x6C, 0x3C, 0x1F, 0xE1, 0x3F, 0xD8, 0xF3, 0xC4, 0x3E, 0x0E, 0xF0, 0x6C, 0xC1, 0x97, 0x41, 0xB6, 0x16, 0x77, 0x5B, 0x3A, 0x2C, 0xA0, 0x29, 0x23, 0xDC, 0xF3, 0x5C, 0xB9, 0xA5, 0x1C, 0x16, 0x1A, 0x50, 0xA5, 0x4E, 0x1A, 0xBD, 0xCC, 0xE3, 0x8D, 0x94, 0xB4, 0x3E, 0x6A, 0xFD, 0x95, 0x24, 0xF8, 0x85, 0xF1, 0x57, 0xC5, 0x17, 0x96, 0x32, 0x78, 0xB6, 0x48, 0xD2, 0xC2, 0xD3, 0xCC, 0x9E, 0x69, 0xB2, 0xC4, 0x2F, 0x7E, 0x7E, 0x95, 0xE9, 0x3A, 0xD7, 0xC1, 0x7F, 0xD9, 0x7B, 0xC5, 0xB7, 0x0D, 0x79, 0xF1, 0x07, 0xE3, 0xD5, 0xD4, 0xB2, 0x23, 0x7C, 0xF0, 0x5A, 0xDD, 0x2A, 0x0D, 0xDD, 0x3A, 0x01, 0xCD, 0x73, 0x5F, 0xF0, 0x4D, 0x9D, 0x26, 0x6D, 0x52, 0x7F, 0x1C, 0x43, 0x68, 0x85, 0xE7, 0x7D, 0x1F, 0xCB, 0x85, 0x49, 0xFB, 0xCC, 0xCA, 0x70, 0x2B, 0x91, 0xF1, 0x9F, 0xEC, 0x2D, 0xFB, 0x47, 0xF8, 0x67, 0x46, 0xBA, 0xF1, 0x8D, 0xD6, 0x95, 0x15, 0xC4, 0x4A, 0xCD, 0x24, 0xD6, 0x90, 0x4D, 0x99, 0x15, 0x72, 0x4F, 0xE3, 0xEB, 0x8A, 0xF3, 0x6A, 0xE5, 0xD8, 0x6A, 0x99, 0x8C, 0x9A, 0x7C, 0xBA, 0x2F, 0xC9, 0x1A, 0xCB, 0x16, 0x9E, 0xEE, 0xC7, 0x0D, 0xE3, 0xDF, 0x05, 0xE8, 0x0F, 0xF1, 0x64, 0xF8, 0x17, 0xE1, 0x06, 0xA1, 0x35, 0xFD, 0x9D, 0xDD, 0xD2, 0xC3, 0xA7, 0xC9, 0x2B, 0x6E, 0x67, 0x27, 0x83, 0x5F, 0x40, 0x58, 0xFE, 0xC9, 0xBF, 0xB2, 0xD7, 0x82, 0x05, 0x97, 0x80, 0xBE, 0x30, 0x78, 0xBA, 0xE2, 0x5F, 0x12, 0xEA, 0x11, 0xAF, 0x99, 0xE5, 0xDD, 0x05, 0x10, 0x39, 0x1C, 0x2E, 0x31, 0xC1, 0xCF, 0xAD, 0x78, 0x6F, 0xEC, 0xC9, 0xAE, 0x69, 0x3E, 0x1C, 0xFD, 0xA2, 0x3C, 0x37, 0xA8, 0xF8, 0x90, 0x79, 0x51, 0xC3, 0xA8, 0x88, 0xE4, 0xF3, 0x97, 0x1E, 0x5B, 0x13, 0x8C, 0x11, 0xEB, 0x9A, 0xF6, 0xAF, 0x8D, 0xDF, 0xB2, 0x27, 0xC6, 0x9F, 0x88, 0x1F, 0xB4, 0xF4, 0xFE, 0x22, 0xD3, 0x22, 0x79, 0x34, 0x2B, 0xEB, 0xF4, 0xBA, 0x4D, 0x53, 0xCC, 0xE2, 0x24, 0xDD, 0x9D, 0xB5, 0xE9, 0xE3, 0xAD, 0x49, 0x46, 0x9B, 0x76, 0x49, 0x6F, 0xDC, 0x8A, 0x98, 0xD7, 0xCD, 0x6B, 0xBB, 0x1E, 0x79, 0xAB, 0x7C, 0x01, 0xB8, 0xF8, 0x23, 0xFB, 0x4B, 0x58, 0x78, 0x29, 0x6F, 0x1A, 0xE2, 0xC2, 0x78, 0xDA, 0x7B, 0x0B, 0xA3, 0xD5, 0xA3, 0xF4, 0xCF, 0xA8, 0xE2, 0x8A, 0xF4, 0xEF, 0x8B, 0xDE, 0x29, 0xD2, 0x7C, 0x61, 0xFB, 0x55, 0xE9, 0x1A, 0x16, 0x8D, 0x28, 0x98, 0x78, 0x7B, 0x48, 0x30, 0x4D, 0x2F, 0xF7, 0x9C, 0x2E, 0x3F, 0xA5, 0x15, 0xF3, 0xB8, 0xE9, 0xB9, 0x55, 0x4E, 0xDD, 0x0D, 0x25, 0x88, 0x9F, 0x2C, 0x6C, 0xFA, 0x75, 0x38, 0x8F, 0x80, 0x5E, 0x35, 0xF8, 0x81, 0xAE, 0x5B, 0xF8, 0x2F, 0x4C, 0xD4, 0x7C, 0x3D, 0x04, 0xBE, 0x1D, 0xD3, 0x3C, 0x5A, 0x21, 0xB6, 0xD5, 0x19, 0x47, 0x99, 0x13, 0x03, 0x91, 0x18, 0x3D, 0x71, 0x5E, 0xCD, 0xA2, 0x7C, 0x04, 0xF8, 0xA7, 0xE1, 0xAF, 0xDA, 0x43, 0xC5, 0x9F, 0x17, 0x35, 0x9D, 0x4E, 0xCC, 0xE8, 0x73, 0x69, 0x57, 0x9F, 0x67, 0x87, 0xED, 0xC3, 0xCD, 0xF9, 0x90, 0xE3, 0x28, 0x4F, 0x1D, 0x2B, 0xCF, 0xFE, 0x08, 0x40, 0x34, 0x4F, 0x84, 0xBF, 0x0B, 0x7C, 0x3B, 0x1C, 0x7B, 0xAE, 0x35, 0xCF, 0x1C, 0x49, 0x7A, 0xEA, 0x7A, 0x84, 0x8C, 0xF5, 0xFA, 0x57, 0x99, 0xFE, 0xD1, 0x3F, 0x16, 0xFC, 0x63, 0x2F, 0xC7, 0x4F, 0x14, 0x0D, 0x13, 0xC4, 0xF7, 0x2B, 0x6C, 0x2F, 0xE5, 0x84, 0x24, 0x53, 0x9D, 0xBB, 0x3A, 0x63, 0xE9, 0x5E, 0xAE, 0x1A, 0x95, 0x7C, 0x65, 0x55, 0x18, 0xF6, 0xFF, 0x00, 0x82, 0x5D, 0x5C, 0x32, 0x82, 0xF7, 0xF6, 0x32, 0xBF, 0x66, 0x3F, 0x0C, 0x37, 0x8B, 0xBF, 0x68, 0xDD, 0x03, 0x4A, 0x30, 0x17, 0x45, 0xD5, 0xDA, 0x67, 0x4C, 0x67, 0xE5, 0x56, 0x24, 0xD7, 0xD9, 0x9E, 0x25, 0xF1, 0xAD, 0xD7, 0xC7, 0x3B, 0x1F, 0x8B, 0x9F, 0x07, 0xFE, 0xD0, 0xAE, 0xBA, 0x3E, 0x9F, 0xFF, 0x00, 0x12, 0xE8, 0x47, 0xF0, 0x90, 0xBE, 0x9E, 0xB9, 0xCD, 0x7C, 0xEB, 0xFF, 0x00, 0x04, 0xD7, 0xF0, 0xF4, 0x5A, 0xE7, 0xC7, 0xCB, 0xCF, 0x16, 0x5C, 0xC9, 0xB6, 0x2D, 0x23, 0x4F, 0x76, 0x67, 0x27, 0x80, 0xCE, 0x48, 0xCE, 0x7B, 0x57, 0xD2, 0x9F, 0x05, 0xFE, 0x14, 0x7C, 0x31, 0xF0, 0x6F, 0xC6, 0xDF, 0x11, 0xF8, 0xC7, 0x49, 0xF8, 0xBF, 0x6B, 0xAA, 0x5F, 0xF8, 0x91, 0x24, 0x5B, 0x8D, 0x24, 0x4E, 0x09, 0x1C, 0x1E, 0x3F, 0x0C, 0xD5, 0xE6, 0x32, 0x9F, 0xD7, 0x23, 0x1F, 0xE5, 0xE5, 0x39, 0xE1, 0x82, 0x9D, 0x58, 0xDE, 0x2B, 0x43, 0xE0, 0xBF, 0x85, 0x5A, 0x07, 0x84, 0x35, 0xFF, 0x00, 0x1C, 0xD9, 0xE8, 0x7F, 0x11, 0x75, 0xC9, 0x74, 0xDD, 0x2F, 0xCC, 0x65, 0xBE, 0xB8, 0x85, 0x7E, 0x68, 0xF1, 0x91, 0x81, 0xF8, 0xD7, 0x43, 0xF0, 0xE7, 0x42, 0xD5, 0x2D, 0xFF, 0x00, 0x69, 0xDD, 0x3F, 0x45, 0xF8, 0x29, 0xAA, 0x49, 0x76, 0xD1, 0x6A, 0xA0, 0xE9, 0x97, 0x52, 0x71, 0xBA, 0x25, 0x3C, 0x96, 0xFC, 0x2B, 0x9D, 0xF8, 0xA3, 0xA0, 0xC1, 0xE1, 0x6F, 0x89, 0xFE, 0x21, 0xD0, 0x15, 0x98, 0x7D, 0x97, 0x55, 0x99, 0x57, 0x8E, 0x70, 0x5B, 0x3F, 0xD6, 0xBD, 0x03, 0xF6, 0x0F, 0xD6, 0xF4, 0x4F, 0x0F, 0x7E, 0xD4, 0xBA, 0x3D, 0xEE, 0xB9, 0x70, 0x90, 0xC2, 0xF6, 0xD2, 0xC5, 0x1C, 0xB2, 0x1C, 0x01, 0x23, 0x0C, 0x2F, 0xEB, 0x5E, 0xED, 0x5E, 0x79, 0x60, 0x9D, 0x4B, 0xFD, 0x94, 0x65, 0x1A, 0x2E, 0x33, 0xE5, 0xEA, 0x7D, 0x55, 0xF1, 0x03, 0x5A, 0xFD, 0x91, 0xA2, 0xF8, 0xD1, 0x61, 0x27, 0xC5, 0x2B, 0x3B, 0x49, 0x7C, 0x5F, 0x06, 0xC4, 0x9A, 0xE6, 0x35, 0xF9, 0x3C, 0xEE, 0x3E, 0xF7, 0x6E, 0xB5, 0xF3, 0xEF, 0xED, 0xB9, 0xA7, 0x78, 0xAE, 0xCF, 0xF6, 0xB0, 0xD0, 0xE5, 0xF1, 0x25, 0xDC, 0x73, 0xD8, 0x5E, 0x5E, 0x5B, 0x3E, 0x91, 0xE4, 0xAE, 0x14, 0x45, 0xB8, 0x0C, 0x0F, 0x7F, 0x7A, 0xCD, 0xF8, 0xC3, 0xFB, 0x2E, 0xFC, 0x7C, 0xD7, 0xBF, 0x69, 0x8B, 0xEB, 0x0B, 0x4F, 0x0C, 0x5C, 0xDC, 0x41, 0x7F, 0xAA, 0x8B, 0x88, 0x35, 0x21, 0xCC, 0x62, 0x2D, 0xC0, 0xE4, 0xB7, 0x6E, 0x2B, 0xA5, 0xFD, 0xB5, 0x7C, 0x4D, 0xA2, 0xEA, 0x5F, 0x1E, 0x3E, 0x1E, 0xF8, 0x32, 0xCA, 0xF5, 0x2E, 0x27, 0xD0, 0x1A, 0xDA, 0xDE, 0xFE, 0x54, 0x6D, 0xC0, 0x3E, 0x57, 0x8C, 0xD7, 0x83, 0x42, 0x87, 0x24, 0xF9, 0x93, 0xE6, 0xD1, 0xDC, 0xD9, 0x61, 0x3F, 0x79, 0xB1, 0x9F, 0xFF, 0x00, 0x05, 0x26, 0xBA, 0x27, 0xE2, 0xDE, 0x8D, 0x6C, 0x98, 0x1E, 0x5E, 0x85, 0x10, 0x03, 0xF0, 0x15, 0xF3, 0xAA, 0x5C, 0xBA, 0xA3, 0x79, 0xA9, 0x8C, 0x72, 0x08, 0x35, 0xF5, 0xD7, 0xED, 0xB3, 0xFB, 0x31, 0xFC, 0x79, 0xF8, 0xBF, 0xF1, 0x5E, 0xD3, 0xC4, 0xFF, 0x00, 0x0D, 0xBC, 0x1C, 0x2F, 0xAC, 0xA3, 0xD2, 0xE1, 0x8C, 0x4C, 0xF3, 0xAA, 0x00, 0x40, 0xF7, 0xAF, 0x05, 0xF8, 0x99, 0xFB, 0x2A, 0x7C, 0x75, 0xF8, 0x45, 0xE1, 0x03, 0xE3, 0x0F, 0x89, 0x1E, 0x1B, 0xB6, 0xB2, 0xB3, 0x79, 0x84, 0x59, 0x5B, 0x85, 0x66, 0x04, 0x8E, 0x38, 0xAF, 0x73, 0x27, 0xC4, 0xE1, 0xE9, 0xE1, 0x54, 0x6F, 0xA9, 0xCB, 0x88, 0xA1, 0x35, 0x2B, 0x24, 0x7A, 0x87, 0xEC, 0x4D, 0xE1, 0xBB, 0x7F, 0x87, 0x5E, 0x06, 0xF1, 0x3F, 0xED, 0x3D, 0xE2, 0x7B, 0x74, 0x55, 0xB3, 0xB4, 0x7B, 0x7D, 0x13, 0xCC, 0x52, 0x59, 0xE4, 0xC7, 0x24, 0x03, 0xEF, 0xDE, 0xBA, 0x0F, 0x06, 0x6B, 0xFA, 0xB7, 0x88, 0x3F, 0xE0, 0x9F, 0x9E, 0x32, 0xF1, 0x46, 0xAC, 0xF2, 0x3D, 0xD6, 0xAB, 0xE2, 0x39, 0xA5, 0x7D, 0xCA, 0x72, 0x73, 0xB7, 0x07, 0x18, 0xCE, 0x2B, 0x37, 0xE1, 0x27, 0xED, 0xB1, 0xF0, 0xD3, 0xC0, 0x5F, 0x04, 0xB4, 0xCF, 0x85, 0x9E, 0x23, 0xF8, 0x63, 0x0E, 0xA6, 0xB6, 0x71, 0x7E, 0xFB, 0xCF, 0xC6, 0xD9, 0x1F, 0x9E, 0x48, 0xAF, 0x58, 0xF1, 0x57, 0xED, 0x05, 0xE0, 0x09, 0xBF, 0x62, 0xFB, 0xAF, 0x88, 0x1A, 0x07, 0x83, 0xB4, 0xDB, 0x38, 0xF5, 0x29, 0xDE, 0xCE, 0x3D, 0x26, 0x3C, 0x05, 0x42, 0x78, 0xDF, 0x8A, 0xF1, 0xB1, 0xB2, 0xAF, 0x2C, 0x47, 0x34, 0xA3, 0xBC, 0x8D, 0xA8, 0xE0, 0xEA, 0xF2, 0x68, 0x7C, 0x79, 0xF0, 0x63, 0xE3, 0x5F, 0xC4, 0x9F, 0x82, 0x97, 0x17, 0xBA, 0xD7, 0xC3, 0x99, 0x66, 0x41, 0x3C, 0x40, 0x5E, 0x34, 0x50, 0x6E, 0x50, 0x00, 0xC0, 0xC9, 0x20, 0xE2, 0xBD, 0x47, 0xF6, 0x4B, 0xFD, 0xA1, 0x3E, 0x2A, 0x78, 0x8B, 0xF6, 0x8E, 0xD3, 0xB4, 0x8D, 0x7F, 0xC4, 0xF3, 0x5F, 0x5A, 0xEB, 0xD2, 0xBA, 0x5E, 0xC1, 0x3C, 0x99, 0x52, 0x31, 0xD4, 0x0E, 0x95, 0x5B, 0xF6, 0x28, 0xF1, 0xA7, 0xC2, 0xE4, 0xB1, 0xF1, 0x27, 0xC2, 0x5F, 0x89, 0x77, 0xB6, 0xF6, 0x70, 0x78, 0x8A, 0x12, 0xB6, 0x9A, 0x84, 0xD8, 0xC4, 0x4C, 0x7B, 0x64, 0xD7, 0x7F, 0xF0, 0xEB, 0xE0, 0xCF, 0xC1, 0x0F, 0xD9, 0x13, 0x5D, 0xBA, 0xF8, 0xB9, 0xE2, 0x3F, 0x8B, 0x76, 0x1A, 0xF5, 0xDD, 0xBC, 0x52, 0x0D, 0x1A, 0xCA, 0xD1, 0xC1, 0x2A, 0x08, 0x38, 0x24, 0x7E, 0x55, 0xD9, 0x8D, 0x9D, 0x19, 0x29, 0xA9, 0x53, 0xF7, 0xEC, 0xAC, 0xFE, 0xE3, 0x48, 0x61, 0x31, 0x13, 0x8F, 0xBA, 0x8E, 0x27, 0xC0, 0xDF, 0x08, 0x34, 0x0F, 0x1A, 0x7E, 0xDE, 0xDA, 0xA7, 0x85, 0xE6, 0xB5, 0x43, 0xA7, 0xE9, 0xD7, 0x72, 0xDF, 0x4B, 0x08, 0x18, 0xDC, 0x50, 0xE4, 0x0A, 0xD7, 0xF1, 0x47, 0xED, 0xE5, 0xF1, 0x2E, 0xC3, 0xE3, 0xBA, 0x69, 0x5A, 0x3C, 0xE5, 0x74, 0x6B, 0x6D, 0x55, 0x2C, 0xE3, 0xD2, 0xC2, 0xF0, 0xD1, 0x87, 0xDB, 0xDB, 0xBD, 0x79, 0xC7, 0xC2, 0x9F, 0xDA, 0x21, 0x3C, 0x13, 0xFB, 0x4C, 0xC9, 0xF1, 0xAE, 0xF9, 0x1D, 0xAD, 0xEF, 0xEE, 0x24, 0x5B, 0xC4, 0x03, 0x9F, 0x2D, 0xCF, 0xFF, 0x00, 0xAA, 0xBD, 0xCE, 0xF7, 0xC3, 0xDF, 0xB0, 0xE6, 0x93, 0xE3, 0x89, 0x3F, 0x68, 0x78, 0x7C, 0x58, 0xF7, 0x77, 0x32, 0xC8, 0x6E, 0x62, 0xD1, 0x14, 0x64, 0x09, 0x8F, 0x3F, 0x77, 0xB7, 0x35, 0xCD, 0x89, 0x8C, 0x94, 0xAD, 0x56, 0x37, 0xF7, 0x55, 0x81, 0x61, 0x25, 0x37, 0x78, 0x74, 0xD0, 0x97, 0xE2, 0x8F, 0x82, 0x34, 0x6D, 0x27, 0xF6, 0xBC, 0x1E, 0x23, 0xD0, 0x6D, 0x02, 0x7F, 0x69, 0x78, 0x70, 0x5C, 0xDD, 0xC5, 0x8E, 0x15, 0xF6, 0x9E, 0xDE, 0xA6, 0x8A, 0xE2, 0x7C, 0x1D, 0xF1, 0x43, 0x5D, 0xF8, 0xA7, 0xF1, 0x7B, 0xC4, 0x5F, 0x15, 0x75, 0x7B, 0x47, 0x82, 0xC9, 0xAC, 0x24, 0x8A, 0xC5, 0x24, 0x3B, 0x76, 0xA0, 0x46, 0xC0, 0x14, 0x57, 0xCC, 0xE3, 0xB9, 0xE3, 0x55, 0x2F, 0x23, 0xD2, 0xA3, 0x96, 0x4E, 0xB4, 0x39, 0x99, 0xCE, 0x7C, 0x1C, 0xF8, 0x9D, 0xE2, 0x18, 0x4E, 0x8C, 0xB2, 0x5B, 0xD9, 0xCC, 0x7C, 0x33, 0xE1, 0xEB, 0x95, 0xD1, 0xDA, 0x68, 0x49, 0x30, 0x96, 0xDD, 0x97, 0xE0, 0x8C, 0xB7, 0xBD, 0x7C, 0xF9, 0xA8, 0xEB, 0xFA, 0xA5, 0xD6, 0xB5, 0x77, 0xA8, 0x5C, 0x5C, 0x97, 0x96, 0x6B, 0x96, 0x79, 0x19, 0xBF, 0x88, 0xEE, 0xEF, 0x45, 0x15, 0xF7, 0x19, 0x42, 0x5B, 0xF9, 0x1D, 0xD9, 0x8A, 0x4E, 0x0F, 0xD4, 0xF6, 0xAF, 0xD9, 0x2B, 0xE2, 0x16, 0xBD, 0xE0, 0x6F, 0x07, 0xF8, 0xBA, 0x5D, 0x0A, 0x3B, 0x61, 0x26, 0xA1, 0x61, 0xB2, 0x69, 0x25, 0x88, 0x96, 0x51, 0xB4, 0xFD, 0xD2, 0x08, 0xC5, 0x70, 0xBF, 0x05, 0x7C, 0x7F, 0xE2, 0xAF, 0x0B, 0xFC, 0x5F, 0xD2, 0xFC, 0x51, 0xA6, 0x6A, 0x6C, 0x6E, 0xA3, 0xBD, 0xF3, 0x01, 0x97, 0x2C, 0xAC, 0x73, 0xDC, 0x77, 0x14, 0x51, 0x58, 0xE2, 0x75, 0xA9, 0x22, 0xE8, 0x36, 0xA8, 0xC2, 0xDE, 0x65, 0xEF, 0x8D, 0x7A, 0x94, 0xDE, 0x31, 0xF8, 0xB7, 0xAB, 0x78, 0x83, 0x53, 0x8A, 0x38, 0xE6, 0xBC, 0x9B, 0x7C, 0xCB, 0x6C, 0x9B, 0x13, 0x77, 0xA8, 0x1C, 0xD6, 0x4E, 0x99, 0xA2, 0xC5, 0x0D, 0xE2, 0xDD, 0xDB, 0xDE, 0xDC, 0xC5, 0x2C, 0x38, 0x78, 0xDE, 0x39, 0x00, 0x2A, 0x47, 0x4E, 0xD4, 0x51, 0x5E, 0x9D, 0x2D, 0x70, 0x4E, 0xE7, 0x87, 0x88, 0x6F, 0xEB, 0x9F, 0x33, 0xE8, 0xCF, 0x0E, 0xFE, 0xD2, 0x1F, 0x1A, 0x9F, 0xE1, 0xE2, 0x68, 0xAF, 0xE3, 0xBB, 0xA3, 0x12, 0xC3, 0xB1, 0x64, 0x21, 0x7C, 0xC0, 0x3A, 0x7D, 0xEC, 0x66, 0xBC, 0x7B, 0xC4, 0x5E, 0x1B, 0xB6, 0x83, 0x57, 0x7F, 0x17, 0x0B, 0xFB, 0xB7, 0xD4, 0x22, 0x9D, 0x67, 0x4B, 0x89, 0x26, 0xDC, 0x43, 0x82, 0x0E, 0x79, 0x1E, 0xB4, 0x51, 0x5E, 0x3D, 0x24, 0x95, 0x37, 0x6F, 0xEB, 0x43, 0xEA, 0x6B, 0xC6, 0x3F, 0x55, 0x83, 0xB7, 0x53, 0x4F, 0x54, 0xFD, 0xB2, 0xFF, 0x00, 0x68, 0xAB, 0x7B, 0x25, 0xB6, 0x8F, 0xE2, 0x15, 0xC0, 0x54, 0xF9, 0x17, 0x8E, 0x40, 0xE9, 0x5C, 0x37, 0x8D, 0x3E, 0x3E, 0xFC, 0x54, 0xF1, 0xDD, 0xBF, 0xF6, 0x47, 0x8B, 0x3C, 0x51, 0x35, 0xED, 0xB0, 0x9B, 0x22, 0x29, 0x89, 0x20, 0x1E, 0x99, 0xC6, 0x68, 0xA2, 0xB2, 0xA7, 0x29, 0x29, 0x2B, 0x3E, 0xA7, 0xAB, 0x56, 0x9D, 0x37, 0x83, 0xD5, 0x23, 0xAF, 0xF0, 0x3F, 0xC3, 0x4F, 0x09, 0xF8, 0x8F, 0x40, 0xFE, 0xD2, 0xD4, 0xAC, 0xDC, 0xC9, 0xE6, 0x11, 0xF2, 0x3E, 0x07, 0xE5, 0x5D, 0x5E, 0x97, 0xF0, 0xBF, 0xC2, 0xD3, 0x69, 0xF1, 0x69, 0x52, 0x8B, 0xA6, 0xB5, 0x0E, 0x4F, 0xD9, 0x8D, 0xC9, 0xD9, 0x9F, 0xA5, 0x14, 0x57, 0x55, 0x76, 0xDE, 0xAC, 0xF9, 0xAA, 0x51, 0x8A, 0x6E, 0xC8, 0xB1, 0xFF, 0x00, 0x0A, 0x73, 0xC0, 0x29, 0x21, 0x31, 0xE9, 0x1B, 0x4A, 0x7D, 0xD2, 0xAF, 0xC8, 0xA8, 0xAE, 0xFE, 0x1E, 0xF8, 0x5D, 0x59, 0x63, 0x7B, 0x13, 0x20, 0x1D, 0x3C, 0xC6, 0xCD, 0x14, 0x57, 0x9F, 0x39, 0x49, 0xBB, 0xB6, 0x7A, 0x18, 0x38, 0x43, 0x95, 0xE8, 0x57, 0x5F, 0x05, 0xF8, 0x5D, 0x18, 0xAA, 0xE8, 0xD0, 0x72, 0x79, 0xFD, 0xD8, 0xAB, 0xD6, 0xBE, 0x13, 0xF0, 0xEC, 0x2C, 0x8C, 0xBA, 0x54, 0x27, 0x9F, 0xBA, 0x50, 0x62, 0x8A, 0x2B, 0x9F, 0x15, 0x52, 0xA7, 0x26, 0xEC, 0xED, 0x8D, 0x1A, 0x4A, 0x4E, 0xD1, 0x5B, 0x76, 0x36, 0x2E, 0x63, 0x87, 0x49, 0xF0, 0xDD, 0xCF, 0xD8, 0x20, 0x44, 0xC5, 0xAC, 0xC0, 0x00, 0x3D, 0x50, 0xD1, 0x45, 0x15, 0xE2, 0x62, 0x3D, 0xE9, 0x26, 0xCF, 0xA8, 0xA3, 0x46, 0x92, 0xA3, 0x1B, 0x45, 0x6D, 0xD8,
		0xFF, 0xD9}
	url, err := PutLocalToCloud(data, int64(len(data)), "foo/bytes.jpg")
	t.Log(url, err)
}

func TestPutLocalFileToCloud(t *testing.T) {
	url, err := PutLocalFileToCloud("E:\\_t\\upup.jpg", "foo/bar.jpg")
	t.Log(url, err)
}

func TestGetCloudFileUrl(t *testing.T) {
	url := GetCloudFileUrl("foo/bar.jpg")
	t.Log(url)
}

func TestPutLocalFileToCloudWithoutKey(t *testing.T) {
	url, err := PutLocalFileToCloudWithoutKey("E:\\_t\\upup.jpg")
	t.Log(url, err)
}
