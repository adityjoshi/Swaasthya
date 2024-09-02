import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Login from './components/login.component';
import SignUp from './components/signup.component';
import PatientForm from './components/patient'; 
import StaffForm from './components/staff';
import Hospital_SignUp from './components/hospital_signup';
import Hospital_Login from './components/hospital_login';
import OtpLogin from './components/loginOtp'; 

function App() {
  return (
    <Router>
      <div className="App">
        <nav className="navbar navbar-expand-lg navbar-light fixed-top">
          <div className="container">
            <Link className="navbar-brand" to="/">
              SWAASTHYA
            </Link>
            <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
              <ul className="navbar-nav ml-auto">
                <li className="nav-item">
                  <Link className="nav-link" to="/sign-in">
                    Login
                  </Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link" to="/sign-up">
                    Sign Up
                  </Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link" to="/hospital-sign-up">
                    Hospital Sign Up
                  </Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link" to="/hospital-login">
                    Hospital Login
                  </Link>
                </li>
                {/* <li className="nav-item">
                  <Link className="nav-link" to="/otp-login">
                    OTP Login
                  </Link>
                </li> */}
              </ul>
            </div>
          </div>
        </nav>

        <div className="auth-wrapper">
          <div className="auth-inner">
            <Routes>
              <Route path="/" element={<Login />} />
              <Route path="/sign-in" element={<Login />} />
              <Route path="/sign-up" element={<SignUp />} />
              <Route path="/patient" element={<PatientForm />} />
              <Route path="/staff" element={<StaffForm />} />
              <Route path="/hospital-sign-up" element={<Hospital_SignUp />} />
              <Route path="/hospital-login" element={<Hospital_Login />} />
              <Route path="/otp-login" element={<OtpLogin />} /> {/* Route for OtpLogin */}
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
