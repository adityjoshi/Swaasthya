import React from "react";
import { useNavigate } from 'react-router-dom';
import "./CommonNavigation.css"

function CommonNavigation() {
    const navigate = useNavigate(); 

    return(
        <header className="navbar">
        <div className="logo">
            <a href="/">
                <img src="/assets/SIH_logo2.png" alt="Logo" width="40" height="92" />
            </a>
        </div>

        <ul className="nav">
            <li><a href="#" className="nav-link">Home</a></li>
            <li><a href="#" className="nav-link">FAQs</a></li>
            <li><a href="#" className="nav-link">About</a></li>
            <li><a href="#hospitals" className="nav-link">Hospitals</a></li>
            <li><a href="#" className="nav-link">Appointments</a></li>
        </ul>

        <div className="auth-buttons">
            <button className="btn login" onClick={() => navigate('./Login')}>
                Login
            </button>
            <button className="btn signup">
                Sign-up
            </button>
        </div>
    </header>
    )
}

export default CommonNavigation;