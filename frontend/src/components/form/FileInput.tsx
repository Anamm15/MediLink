// components/ui/ModernFileInput.tsx
'use client';

import { useState, useRef, useEffect, DragEvent } from 'react';
// npm install react-icons
import { FiUploadCloud, FiFile, FiX } from 'react-icons/fi';
import Image from 'next/image';

// Helper function untuk memformat ukuran file
const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
};

type ModernFileInputProps = {
  // Callback untuk mengirim file terpilih ke parent component
  onChange: (files: File[]) => void;
  // Prop opsional untuk menerima file dari parent (jika diperlukan)
  value?: File[];
};

export default function ModernFileInput({ onChange, value = [] }: ModernFileInputProps) {
  const [selectedFiles, setSelectedFiles] = useState<File[]>(value);
  const [isDragging, setIsDragging] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);
  
  // State untuk menyimpan URL preview gambar
  const [imagePreviews, setImagePreviews] = useState<{[key: string]: string}>({});

  // Sinkronisasi state internal dengan prop 'value'
  useEffect(() => {
    setSelectedFiles(value);
  }, [value]);

  // Membuat dan membersihkan URL preview gambar
  useEffect(() => {
    const newImagePreviews: {[key: string]: string} = {};
    selectedFiles.forEach(file => {
      if (file.type.startsWith('image/')) {
        newImagePreviews[file.name] = URL.createObjectURL(file);
      }
    });
    setImagePreviews(newImagePreviews);

    // Cleanup function untuk mencegah memory leak
    return () => {
      Object.values(newImagePreviews).forEach(url => URL.revokeObjectURL(url));
    };
  }, [selectedFiles]);
  
  const handleFileChange = (files: FileList | null) => {
    if (files) {
      const newFiles = Array.from(files);
      const updatedFiles = [...selectedFiles, ...newFiles];
      setSelectedFiles(updatedFiles);
      onChange(updatedFiles);
    }
  };

  const handleDragEnter = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(false);
  };

  const handleDragOver = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault(); // Ini penting agar event onDrop bisa berjalan
  };

  const handleDrop = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(false);
    const files = e.dataTransfer.files;
    handleFileChange(files);
  };

  const handleRemoveFile = (indexToRemove: number) => {
    const updatedFiles = selectedFiles.filter((_, index) => index !== indexToRemove);
    setSelectedFiles(updatedFiles);
    onChange(updatedFiles);
  };
  
  const openFileDialog = () => {
    fileInputRef.current?.click();
  };

  return (
    <div className="w-full">
      {/* Area Dropzone */}
      <div
        onClick={openFileDialog}
        onDragEnter={handleDragEnter}
        onDragLeave={handleDragLeave}
        onDragOver={handleDragOver}
        onDrop={handleDrop}
        className={`flex w-full cursor-pointer flex-col items-center justify-center rounded-lg border-2 border-dashed p-8 text-center transition-colors duration-300 ${
          isDragging
            ? 'border-indigo-600 bg-indigo-50'
            : 'border-gray-300 bg-gray-50 hover:bg-gray-100'
        }`}
      >
        <input
          ref={fileInputRef}
          type="file"
          multiple
          onChange={(e) => handleFileChange(e.target.files)}
          className="hidden"
        />
        <FiUploadCloud className={`mb-4 h-12 w-12 transition-colors duration-300 ${isDragging ? 'text-indigo-600' : 'text-gray-400'}`} />
        <p className="font-semibold text-gray-600">
          <span className="text-indigo-600">Klik untuk upload</span> atau seret & lepas file
        </p>
        <p className="text-xs text-gray-500">PNG, JPG, GIF, atau PDF (maks. 10MB)</p>
      </div>

      {/* Daftar File yang Dipilih */}
      {selectedFiles.length > 0 && (
        <div className="mt-6 space-y-3">
          <h3 className="font-semibold text-gray-800">File yang akan diupload:</h3>
          {selectedFiles.map((file, index) => (
            <div
              key={index}
              className="flex items-center justify-between rounded-lg border border-gray-200 bg-white p-3 shadow-sm"
            >
              <div className="flex items-center gap-4">
                {file.type.startsWith('image/') ? (
                   <Image
                      src={imagePreviews[file.name]}
                      alt={file.name}
                      width={40}
                      height={40}
                      className="h-10 w-10 rounded-md object-cover"
                    />
                ) : (
                  <div className="flex h-10 w-10 items-center justify-center rounded-md bg-gray-100">
                    <FiFile className="h-5 w-5 text-gray-500" />
                  </div>
                )}
                <div className="flex flex-col">
                  <span className="text-sm font-medium text-gray-800">{file.name}</span>
                  <span className="text-xs text-gray-500">{formatFileSize(file.size)}</span>
                </div>
              </div>
              <button
                type="button"
                onClick={() => handleRemoveFile(index)}
                className="text-gray-500 hover:text-red-600"
                aria-label={`Hapus file ${file.name}`}
              >
                <FiX className="h-5 w-5" />
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}