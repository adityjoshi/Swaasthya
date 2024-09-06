import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Login from './pages/Login';
import SignUp from './pages/Signup';
import LandingPage from './pages/LandingPage';
import PatientForm from './components/patient'; 
import StaffForm from './components/staff';
import Hospital_Login from './pages/Hospital_Login';
import Hospital_SignUp from './pages/Hospital_Signup';
import Patient_Dashboard from './pages/PatientDashboard';
import OtpLogin from './pages/loginOtp'; 

function App() {
  return (
    <Router>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/patient" element={<PatientForm />} />
          <Route path="/staff" element={<StaffForm />} />
          <Route path="/hospital-login" element={<Hospital_Login />} />
          <Route path="/hospital-signup" element={<Hospital_SignUp />} />
          <Route path="/patient-dashboard" element={<Patient_Dashboard />} />
          <Route path="/otp-login" element={<OtpLogin />} /> {/* Route for OtpLogin */}
        </Routes>
    </Router>
  );
}

export default App;
