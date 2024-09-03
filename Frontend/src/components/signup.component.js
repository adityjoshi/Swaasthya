import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './signUp.css'; 

const SignUp = () => {
  const [userType, setUserType] = useState("");
  const navigate = useNavigate();

  const handleUserTypeChange = (e) => {
    setUserType(e.target.value);
    if (e.target.value === "Patient") {
      navigate('/patient');
    } else if (e.target.value === "Staff") {
      navigate('/staff'); 
    }
  };

  return (
    <div className="signup-container">
      <form className="signup-form">
        <img src="../assets/Logo1.png" alt="Logo" className="form-logo" />
        <h3>Sign Up</h3>

        <div className="mb-3">
          <label>Full Name</label>
          <input type="text" className="form-control" placeholder="Full name" />
        </div>

        <div className="mb-3">
          <label>Gender</label>
          <select className="form-control" id="mf">
            <option value="" selected>Choose...</option>
            <option value="Male">Male</option>
            <option value="Female">Female</option>
          </select>
        </div>

        <div className="mb-3">
          <label>Contact Number</label>
          <input type="number" className="form-control" placeholder="Contact Number" />
        </div>

        <div className="mb-3">
          <label>Email</label>
          <input type="email" className="form-control" placeholder="Enter email" />
        </div>

        <div className="mb-3">
          <label>Password</label>
          <input type="password" className="form-control" placeholder="Enter password" />
        </div>

        <div className="mb-3">
          <label>Language</label>
          <select className="form-control" id="lang">
            <option value="" selected>Choose...</option>
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
          >
            <option value="" selected>Choose...</option>
            <option value="Patient">Patient</option>
            <option value="Staff">Staff</option>
          </select>
        </div>

        <div className="d-grid">
          <button type="button" className="btn btn-primary">
            Sign Up
          </button>
        </div>

        <p className="already-registered text-right">
          Already registered <a href="/sign-in">sign in?</a>
        </p>
      </form>
    </div>
  );
};

export default SignUp;
