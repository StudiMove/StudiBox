import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ProfileImageState {
  profileImage: string | null; // Permettre à l'image de profil d'être null
}

const initialState: ProfileImageState = {
  profileImage: null, // Initialiser avec null
};

const profileImageSlice = createSlice({
  name: 'profileImage',
  initialState,
  reducers: {
    setProfileImage(state, action: PayloadAction<string | null>) {
      state.profileImage = action.payload; // Mettre à jour l'image de profil
    },
    clearProfileImage(state) {
      state.profileImage = null; // Réinitialiser l'image de profil
    },
  },
});

export const { setProfileImage, clearProfileImage } = profileImageSlice.actions;
export default profileImageSlice.reducer;
