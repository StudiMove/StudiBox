import React, { useState, useEffect, useRef } from 'react';
import InputField from '../../UI/Inputs/InputField/InputField';

import { useDispatch } from 'react-redux';
import { setProfileImage } from '../../../../store/slices/profileImage.slice';
import { ProfileForm } from '../../../types/api/profile/ProfilForm/ProfilForm';
import {
  ProfileUpdateData,
  TargetProfileUpdateData,
} from '../../../types/api/profile/ProfileRequest';
import {
  getProfileData,
  getProfileDataById,
  updateProfileData,
  updateTargetProfileData,
  updateProfileImage,
  uploadTargetProfileImage,
} from '../../../services/ProfilService';
import { useSelector } from 'react-redux';
import { RootState } from '../../../../store';
import { hasAllAdminPermission } from '../../../config/permissions';
import IconTitle from '../../UI/Typography/IconTitle';
import Container from '../../UI/Container/Container';
import buildingOfficeIcon from '../../Icons/buildingOfficeIcon.svg';
import mapIcon from '../../Icons/mapIcon.svg';
import photoIcon from '../../Icons/photoIcon.svg';
import userIcon from '../../Icons/userIcon.svg';
import PrimaryButton from '../../UI/Button/PrimaryButton';
import SecondaryButtonOutline from '../../UI/Button/SecondaryButtonOutline';
import SaveButton from '../../UI/Button/SaveButton';
import InputFieldList from '../../UI/Inputs/InputFieldList/InputFieldList';

interface ProfilFormComponentProps {
  targetId?: number;
}
const ProfilFormComponent = ({ targetId }: ProfilFormComponentProps) => {
  const [formData, setFormData] = useState<ProfileForm | null>(null);
  const [loading, setLoading] = useState(true);
  const [selectedImage, setSelectedImage] = useState<File | null>(null);
  const [imageName, setImageName] = useState<string>('');
  const role = useSelector((state: RootState) => state.auth.role);
  const isEditable = hasAllAdminPermission(role);
  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const dispatch = useDispatch();
  const profileImage = useSelector(
    (state: RootState) => state.profileImage.profileImage,
  );

  const fetchProfileData = async () => {
    try {
      const data = targetId
        ? await getProfileDataById(targetId)
        : await getProfileData();
      const mappedData: ProfileForm = {
        userId: data.userId,
        email: data.email,
        phone: data.phone,
        profileImage: data.profileImage,
        roles: data.roles,
        organisationName: data.organisationName,
        organisationAddress: data.organisationAddress,
        organisationCity: data.organisationCity,
        organisationPostcode: data.organisationPostcode,
        organisationCountry: data.organisationCountry,
        organisationRegion: data.organisationRegion,
        organisationDescription: data.organisationDescription,
        status: data.status,
        siret: data.siret,
        type: data.type,
        is_validated: data.is_validated,
        is_pending: data.is_pending,
        is_activated: data.is_activated,
      };
      setFormData(mappedData);
      setLoading(false);
    } catch (error) {
      console.error('Erreur lors du chargement des données du profil:', error);
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchProfileData();
  }, []);

  const handleChange = (key: keyof ProfileForm) => (value: string | number) => {
    setFormData((prevData) => {
      if (!prevData) return prevData;

      const updatedData = { ...prevData, [key]: value };

      if (key === 'status') {
        if (value === 'isPending') {
          updatedData.is_validated = false;
          updatedData.is_pending = true;
          updatedData.is_activated = false;
        } else if (value === 'Validate') {
          updatedData.is_validated = true;
          updatedData.is_pending = false;
          updatedData.is_activated = true;
        } else if (value === 'Invalide') {
          updatedData.is_validated = false;
          updatedData.is_pending = false;
          updatedData.is_activated = true;
        } else if (value === 'Suspended') {
          updatedData.is_validated = false;
          updatedData.is_pending = false;
          updatedData.is_activated = false;
        }
      }

      return updatedData;
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (formData) {
      try {
        if (selectedImage) {
          const imageFormData = new FormData();
          imageFormData.append('file', selectedImage);

          let imageUrl: string | undefined;
          if (targetId) {
            // Utilisation de uploadTargetProfileImage si targetId est présent
            imageUrl = await uploadTargetProfileImage(targetId, imageFormData);
          } else {
            // Utilisation de updateProfileImage si pas de targetId
            imageUrl = await updateProfileImage(imageFormData);
          }

          if (imageUrl) {
            formData.profileImage = imageUrl; // Mise à jour de l'image dans formData
            dispatch(setProfileImage(imageUrl)); // Mettre à jour l'image dans Redux
          }
        }

        if (targetId) {
          const updateData: TargetProfileUpdateData = {
            targetId: targetId,
            updateData: {
              name: formData.organisationName,
              address: formData.organisationAddress,
              city: formData.organisationCity,
              postcode: formData.organisationPostcode,
              region: formData.organisationRegion,
              phone: formData.phone,
              country: formData.organisationCountry,
              email: formData.email,
              siret: formData.siret,
              description: formData.organisationDescription,
              status: formData.status,
              type: formData.type,
              is_validated: formData.is_validated,
              is_pending: formData.is_pending,
              is_activated: formData.is_activated,
              profileImage: formData.profileImage,
            },
          };
          await updateTargetProfileData(updateData);
        } else {
          const updateData: ProfileUpdateData = {
            name: formData.organisationName,
            address: formData.organisationAddress,
            city: formData.organisationCity,
            postcode: formData.organisationPostcode,
            region: formData.organisationRegion,
            phone: formData.phone,
            country: formData.organisationCountry,
            email: formData.email,
            siret: formData.siret,
            description: formData.organisationDescription,
            status: formData.status,
            type: formData.type,
            is_validated: formData.is_validated,
            is_pending: formData.is_pending,
            is_activated: formData.is_activated,
            profileImage: formData.profileImage,
          };
          await updateProfileData(updateData);
        }
      } catch (error) {
        console.error(
          'Erreur lors de la mise à jour des données du profil:',
          error,
        );
      }
    }
  };

  const handleChangePictureClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };
  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setSelectedImage(e.target.files[0]);
      setImageName(e.target.files[0].name);
    }
  };

  const handleDeletePicture = () => {
    console.log('Image supprimée');
    setSelectedImage(null);
    setImageName('');
  };

  const handleImageUpload = async () => {
    if (selectedImage) {
      const formData = new FormData();
      formData.append('file', selectedImage);
      try {
        await updateProfileImage(formData);
      } catch (error) {
        console.error(
          "Erreur lors de la mise à jour de l'image de profil:",
          error,
        );
      }
    }
  };

  if (loading) {
    return <div>Chargement...</div>;
  }

  if (!formData) {
    return <div>Erreur lors de la récupération des données du profil.</div>;
  }

  return (
    <div className="pt-8 relative mb-16">
      <IconTitle title="Photo de profil" image={photoIcon} className="mb-6" />
      <div className="flex items-center mb-8">
        <img
          className="rounded-full w-24 h-24 object-contain"
          src={profileImage || ''}
          alt="Photo de profil"
        />
        <div className="flex flex-col ml-6 md:flex-row md:ml-8 gap-3">
          <PrimaryButton
            text="Change picture"
            onClick={handleChangePictureClick}
          />
          <SecondaryButtonOutline
            label={'Delete picture'}
            onClick={handleDeletePicture}
          />
          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            onChange={handleImageChange}
            style={{ display: 'none' }}
          />
          <PrimaryButton
            text="Upload"
            onClick={handleImageUpload}
            className="!bg-green"
          />
        </div>
      </div>
      {imageName && <p>Image sélectionnée : {imageName}</p>}
      <IconTitle className="mb-3" title="Gestion" image={buildingOfficeIcon} />
      <form onSubmit={handleSubmit}>
        {/* Affichage conditionnel basé sur le rôle */}
        {isEditable ? (
          <div className="mb-6">
            <InputFieldList
              label="Statut"
              value={formData?.status || 'isPending'} // Définit la valeur par défaut
              options={[
                { value: 'isPending', label: 'En attente' },
                { value: 'Validate', label: 'Valider' },
                { value: 'Suspended', label: 'Suspendre' },
                { value: 'Invalide', label: 'Invalider' },
              ]}
              onChange={handleChange('status')}
              isEditable={isEditable} // Le champ est désactivé si `isEditable` est false
            />
          </div>
        ) : (
          <InputField
            type="text"
            label="Statut"
            value={formData.status}
            onChange={handleChange('status')}
            size="small"
            isEditable={isEditable}
          />
        )}

        <IconTitle
          className="mb-3"
          title="Informations de l’entreprise"
          image={userIcon}
        />
        <Container variant="two-input-row">
          <InputField
            type="text"
            label="Nom de l’entreprise"
            value={formData.organisationName}
            onChange={handleChange('organisationName')}
            size="small"
            isEditable={isEditable}
          />
          <InputField
            type="text"
            label="Type"
            value={formData.type}
            onChange={handleChange('type')}
            size="small"
            isEditable={isEditable}
          />
        </Container>
        <Container variant="two-input-row">
          <InputField
            type="email"
            label="Email"
            value={formData.email}
            onChange={handleChange('email')}
            size="small"
            isEditable={isEditable}
          />
          <InputField
            type="password"
            label="Mot de Passe"
            value={formData.password}
            onChange={handleChange('password')}
            hasIcon={true}
            size="small"
            isEditable={isEditable}
          />
        </Container>
        <Container variant="two-input-row">
          <InputField
            type="text"
            label="Siret"
            value={formData.siret}
            onChange={handleChange('siret')}
            size="small"
            isEditable={isEditable}
          />
          <InputField
            type="text"
            label="Numéro de téléphone"
            value={formData.phone}
            onChange={handleChange('phone')}
            size="small"
            isEditable={isEditable}
          />
        </Container>
        <InputField
          type="text"
          label="Description"
          value={formData.organisationDescription}
          onChange={handleChange('organisationDescription')}
          size="small"
        />
        <IconTitle className="mb-3" title="Localisation" image={mapIcon} />
        <Container variant="two-input-row">
          <InputField
            type="text"
            label="Ville"
            value={formData.organisationCity}
            onChange={handleChange('organisationCity')}
            size="small"
            isEditable={isEditable}
          />
          <InputField
            type="text"
            label="Région"
            value={formData.organisationRegion}
            onChange={handleChange('organisationRegion')}
            size="small"
            isEditable={isEditable}
          />
        </Container>
        <Container variant="two-input-row">
          <InputField
            type="text"
            label="Code Postal"
            value={formData.organisationPostcode}
            onChange={handleChange('organisationPostcode')}
            size="small"
            isEditable={isEditable}
          />
          <InputField
            type="text"
            label="Pays"
            value={formData.organisationCountry}
            onChange={handleChange('organisationCountry')}
            size="small"
            isEditable={isEditable}
          />
        </Container>
        <Container variant="two-input-row-1">
          <InputField
            type="text"
            label="Adresse"
            value={formData.organisationAddress}
            onChange={handleChange('organisationAddress')}
            size="small"
            isEditable={isEditable}
          />
        </Container>
        <div className="fixed right-4 bottom-4">
          <SaveButton onClick={handleSubmit} />
        </div>
      </form>
    </div>
  );
};

export default ProfilFormComponent;
