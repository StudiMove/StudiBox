import PrimaryButtonOutline from '../../UI/Button/PrimaryButtonOutline';
import trashIcon from '../../Icons/trash/red-trash.svg';
import photoIcon from '../../Icons/photoIcon.svg';
import IconTitle from '../../UI/Typography/IconTitle';
import InputField from '../../UI/Inputs/InputField/InputField';
import H3 from '../../UI/Typography/H3';

interface ImageUploadProps {
  images: File[];
  video_url: string;
  updateVideoUrl: (url: string) => void;
  updateField: (images: File[]) => void;
  uploadImages: (images: File[]) => Promise<string[]>;
}

const MAX_FILES = 4;

const ImageUpload = ({
  images,
  updateField,
  uploadImages,
  video_url,
  updateVideoUrl,
}: ImageUploadProps) => {
  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const newFiles = Array.from(e.target.files);

      if (images.length + newFiles.length > MAX_FILES) {
        alert(`Vous ne pouvez pas télécharger plus de ${MAX_FILES} fichiers.`);
        return;
      }

      const uploadedUrls = await uploadImages(newFiles);
      updateField([...images, ...newFiles]);
      console.log('Uploaded image URLs:', uploadedUrls);
    }
  };

  const handleRemoveFile = (index: number) => {
    const updatedImages = images.filter((_, i) => i !== index);
    updateField(updatedImages);
  };

  return (
    <div>
      <IconTitle title="Images et vidéos" image={photoIcon} className="mb-6" />

      <InputField
        type="text"
        label="Lien vidéo"
        value={video_url}
        onChange={(value: string | number) => updateVideoUrl(String(value))}
        size="small"
        placeholder="URL de la vidéo"
      />

      <H3>Upload une image</H3>
      <div
        className="mb-6 flex flex-col items-center justify-center border-2 border-dashed border-lightGray rounded-lg p-4 h-48"
        onDrop={(e) => {
          e.preventDefault();
          handleFileChange(e as unknown as React.ChangeEvent<HTMLInputElement>);
        }}
      >
        <p className="text-gray-500">Déposer les fichiers ici</p>
        <p className="text-gray-500">ou</p>
        <PrimaryButtonOutline
          label="Sélectionner des fichiers"
          onClick={() => document.getElementById('fileInput')?.click()}
        />
        <input
          id="fileInput"
          type="file"
          multiple
          onChange={handleFileChange}
          className="hidden"
        />
      </div>

      <div>
        {images.map((file: File, index: number) => (
          <div key={index} className="flex items-center justify-between mb-2">
            <p>{file.name}</p>
            <img
              src={trashIcon}
              alt="Supprimer"
              className="h-6 w-6 cursor-pointer"
              onClick={() => handleRemoveFile(index)}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ImageUpload;
