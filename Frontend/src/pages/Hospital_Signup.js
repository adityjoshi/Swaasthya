import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import NavigationLog from '../components/NavigationLog';
// import './signUp.css';

const Hospital_SignUp = () => {
  const [department, setDepartment] = useState("");
  const [otherDepartment, setOtherDepartment] = useState("");
  const navigate = useNavigate();

  const handleDepartmentChange = (e) => {
    setDepartment(e.target.value);
    if (e.target.value === "other") {
      setOtherDepartment(""); 
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    // Add sign-up logic here

    // On successful sign-up, navigate to another page
    navigate('/hospital-login'); // Or any other route
  };

  return (
    <>
      <NavigationLog />
      <div className="signup-container">
        <form className="signup-form" onSubmit={handleSubmit}>
          <img src="/assets/SIH_logo2.png" alt="Logo" className="form-logo" />
          <h3>Sign Up</h3>
          <div className="mb-3">
            <label>Name</label>
            <input type="text" className="form-control" placeholder="Enter your Name" required />
          </div>

          <div className="mb-3">
            <label>Hospital ID</label>
            <input type="number" className="form-control" placeholder="Enter your Hospital ID" required />
          </div>

          <div className="mb-3">
            <label>Address</label>
            <input type="text" className="form-control" placeholder="Enter your Address" required />
          </div>

          <div className="mb-3">
            <label>Email</label>
            <input type="email" className="form-control" placeholder="Enter email" required />
          </div>

          <div className="mb-3">
            <label>Website</label>
            <input type="text" className="form-control" placeholder="Website (If any)" />
          </div>

          <div className="mb-3">
            <label>Department</label>
            <select 
              className="form-control" 
              id="department" 
              value={department}
              onChange={handleDepartmentChange}
              required
            >
              <option value="">Choose...</option>
              <option value="Orthopaedic">Orthopaedic</option>
              <option value="Cardiac">Cardiac</option>
              <option value="Physician">Physician</option>
              <option value="Pediatrician">Pediatrician</option>
              <option value="ENT">ENT</option>
              <option value="Dermatologist">Dermatologist</option>
              <option value="Gynecologist">Gynecologist</option>
              <option value="Dentist">Dentist</option>
              <option value="other">Other</option>
            </select>
            {department === "other" && (
              <div className="mt-3">
                <label>Specify Department</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Specify your department"
                  value={otherDepartment}
                  onChange={(e) => setOtherDepartment(e.target.value)}
                  required
                />
              </div>
            )}
          </div>

          <div className="d-grid">
            <button type="submit" className="submit-button">
              Sign Up
            </button>
          </div>

          <p className="forgot-password text-right">
            Already registered <a href="/hospital-login">sign in?</a>
          </p>
        </form>
      </div>
    </>
  );
};

export default Hospital_SignUp;
