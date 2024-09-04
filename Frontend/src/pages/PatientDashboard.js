import React, { useEffect, useState } from 'react';
import './PatientDashboard.css'; // Import the CSS file

const PatientDashboard = () => {
  const [patientData, setPatientData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Fetch patient data from the backend
    const fetchPatientData = async () => {
      try {
        const response = await fetch('/api/patient-data'); // Replace with your actual API endpoint
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setPatientData(data);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };

    fetchPatientData();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <div className="dashboard-container">
      <div className="patient-dashboard">
        <h2>Welcome, {patientData.patientUsername}</h2>
        <div className="patient-details-inline">
          <span><strong>Patient ID:</strong> {patientData.patientId}</span>
          <span><strong>State:</strong> {patientData.state}</span>
          <span><strong>Bin Code:</strong> {patientData.binCode}</span>
          <span><strong>Aadhar:</strong> {patientData.adhar}</span>
          <span><strong>Medical History:</strong> {patientData.medicalHistory}</span>
          <span><strong>Other Medical Details:</strong> {patientData.otherMedical}</span>
        </div>
      </div>

      {/* Navigation Menu Outside and Below Patient Details */}
      <nav className="patient-navigation">
        <ul>
          <li><a href="#hospitals">Hospitals</a></li>
          <li><a href="#appointments">Appointments</a></li>
          <li><a href="#booking">Booking</a></li>
        </ul>
      </nav>
    </div>
  );
};

export default PatientDashboard;
