import React, { useState } from 'react';
import './StaffForm.css'; 

const StaffForm = () => {
  return (
    <div className="staff-form-container">
      <form className="staff-form">
        <img src="../assets/Logo1.png" alt="Logo" className="form-logo" />
        <h3>STAFF REGISTRATION</h3>

        <div className="mb-3">
          <label>Employee ID</label>
          <input
            type="number"
            className="form-control"
            placeholder="Enter your ID..."
          />
        </div>

        <div className="mb-3">
          <label>Hospital ID</label>
          <input
            type="number"
            className="form-control"
            placeholder="Enter your Hospital ID..."
          />
        </div>

        <div className="mb-3">
          <label>Aadhar</label>
          <input
            type="number"
            className="form-control"
            placeholder="Enter your Aadhar Number..."
          />
        </div>

        <div className="mb-3">
          <label>Designation</label>
          <select className="form-control" id="designation">
            <option value="">Choose...</option>
            <option value="Compounder">Compounder</option>
            <option value="CMO">CMO</option>
            <option value="Doctor">Doctor</option>
            <option value="Manager">Manager</option>
          </select>
        </div>

        <div className="d-grid">
          <button type="submit" className="btn btn-primary">
            Sign Up
          </button>
        </div>
      </form>
    </div>
  );
};

export default StaffForm;
