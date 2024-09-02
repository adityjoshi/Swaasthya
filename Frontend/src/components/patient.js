import React, { useState } from 'react';
import './PatientForm.css'; // Ensure you create a CSS file for styling

const PatientForm = () => {
  const [medicalHistory, setMedicalHistory] = useState("");
  const [otherMedicalHistory, setOtherMedicalHistory] = useState("");

  const handleMedicalHistoryChange = (e) => {
    setMedicalHistory(e.target.value);
    if (e.target.value !== "Other") {
      setOtherMedicalHistory("");
    }
  };

  return (
    <div className="patient-form-container">
      <form className="patient-form">
        <img src="../assets/Logo1.png" alt="Logo" className="form-logo" />
        <h3>PATIENT REGISTRATION</h3>

        <div className="mb-3">
          <label>Username</label>
          <input
            type="text"
            className="form-control"
            placeholder="Username"
          />
        </div>

        <div className="mb-3">
          <label>City</label>
          <input
            type="text"
            className="form-control"
            placeholder="City"
          />
        </div>

        <div className="mb-3">
          <label>State</label>
          <input
            type="text"
            className="form-control"
            placeholder="State"
          />
        </div>

        <div className="mb-3">
          <label>Pin Code</label>
          <input
            type="number"
            className="form-control"
            placeholder="Pin Code"
          />
        </div>

        <div className="mb-3">
          <label>Aadhar</label>
          <input
            type="number"
            className="form-control"
            placeholder="Aadhar Number"
          />
        </div>

        <div className="mb-3">
          <label>Medical History</label>
          <select
            id="medical"
            className="form-control"
            value={medicalHistory}
            onChange={handleMedicalHistoryChange}
          >
            <option value="">Choose...</option>
            <option value="Diabetes">Diabetes</option>
            <option value="Thyroid">Thyroid</option>
            <option value="Cardiac">Cardiac</option>
            <option value="Cancer">Cancer</option>
            <option value="None">None</option>
            <option value="Other">Other</option>
          </select>
        </div>

        {medicalHistory === "Other" && (
          <div className="mb-3">
            <label>Other Medical Issues</label>
            <textarea
              className="form-control"
              placeholder="Describe any other medical issues"
              rows="3"
              value={otherMedicalHistory}
              onChange={(e) => setOtherMedicalHistory(e.target.value)}
            />
          </div>
        )}

        <div className="d-grid">
          <button type="submit" className="btn btn-primary">
            Sign Up
          </button>
        </div>
      </form>
    </div>
  );
};

export default PatientForm;
