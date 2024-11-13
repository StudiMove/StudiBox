import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import store, { persistor } from '../store';
import LoginPage from './pages/LoginPage';
import Layout from './components/UI/Layout/Layout';
import './styles/App.css';
import './styles/tailwind.css';
import ProfilPage from './pages/ProfilPage';
import OrganisationPage from './pages/OrganisationPage';
import SignupPage from './pages/SignupPage';
import EventFormComponent from './components/Form/EventForm/EventFormComponent';
import TargetOrganisationPage from './pages/TargetOrganisationPage';
import EventPage from './pages/Event/EventPage';
import UserPage from './pages/UserPage';
import DashboardPage from './pages/DashboardPage';
import ForgotPasswordPage from './pages/ForgotPasswordPage';
import NotFoundPage from './pages/NotFoundPage';

import ProfilFormComponent from './components/Form/ProfilForm/ProfileFormComponent';
const App = () => {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <Router>
          <Routes>
            {/* Route publique (login) */}
            <Route path="/" element={<SignupPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/forgotpassword" element={<ForgotPasswordPage />} />
            <Route path="*" element={<NotFoundPage />} />
            <Route path="/profile" element={<ProfilFormComponent />} />
            {/* Routes protégées */}
            <Route path="/" element={<Layout />}>
              <Route path="*" element={<NotFoundPage />} />

              <Route path="dashboard" element={<DashboardPage />} />
              <Route path="profil" element={<ProfilPage />} />
              <Route path="organisation" element={<OrganisationPage />} />
              <Route path="events" element={<EventPage />} />
              <Route
                path="/events/createNewEvent"
                element={<EventFormComponent isUpdate={false} />}
              />
              <Route
                path="/organisation/:companyName"
                element={<TargetOrganisationPage />}
              />

              <Route path="user" element={<UserPage />} />
            </Route>
          </Routes>
        </Router>
      </PersistGate>
    </Provider>
  );
};

export default App;
