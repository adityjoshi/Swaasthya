import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import NavigationLog from '../components/NavigationLog';
import './Signup.css';

const SignUp = () => {
  const [userType, setUserType] = useState("");
  const navigate = useNavigate();

  const handleUserTypeChange = (e) => {
    setUserType(e.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (userType === "Patient") {
      navigate('/patient');
    } else if (userType === "Staff") {
      navigate('/staff'); 
    }
  };

  return (
    <>
      <NavigationLog />
      <div className="signup-container">
        <form className="signup-form" onSubmit={handleSubmit}>
          <img src="/assets/SIH_logo2.png" alt="Logo" className="form-logo" />
          <h3>Sign Up</h3>

          <div className="mb-3">
            <label>Full Name</label>
            <input type="text" className="form-control" placeholder="Full name" required />
          </div>

          <div className="mb-3">
            <label>Gender</label>
            <select className="form-control" id="mf">
              <option value="">Choose...</option>
              <option value="Male">Male</option>
              <option value="Female">Female</option>
            </select>
          </div>

          <div className="mb-3">
            <label>Contact Number</label>
            <input type="tel" className="form-control" placeholder="Contact Number" required />
          </div>

          <div className="mb-3">
            <label>Email</label>
            <input type="email" className="form-control" placeholder="Enter email" required />
          </div>

          <div className="mb-3">
            <label>Password</label>
            <input type="password" className="form-control" placeholder="Enter password" required />
          </div>

          <div className="mb-3">
            <label>Language</label>
            <select className="form-control" id="lang">
              <option value="">Choose...</option>
              <option value="Hindi">Hindi</option>
              <option value="English">English</option>
            </select>
          </div>

          <div className="mb-3">
            <label>User Type</label>
            <select 
              className="form-control" 
              id="usertype" 
              value={userType} 
              onChange={handleUserTypeChange}
              required
            >
              <option value="">Choose...</option>
              <option value="Patient">Patient</option>
              <option value="Staff">Staff</option>
            </select>
          </div>

          <div className="d-grid">
            <button type="submit" className="submit-button">
              Sign Up
            </button>
          </div>

          <p className="already-registered text-right">
            Already registered <a href="/sign-in">sign in?</a>
          </p>
        </form>
      </div>
    </>
  );
};

export default SignUp;
