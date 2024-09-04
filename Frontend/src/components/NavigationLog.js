import React from 'react';
import { Link } from 'react-router-dom';
import './NavigationLog.css';

const NavigationLog = () => {
    return (
        <div className="App">
            <nav className="navbar">
                <div className="logo">
                    <Link to="/">
                        <img src="/assets/SIH_logo2.png" alt="SWAASTHYA" />
                    </Link>
                </div>
                <ul className="nav">
                    <li>
                        <Link className="nav-link" to="/login">
                            Login
                        </Link>
                    </li>
                    <li>
                        <Link className="nav-link" to="/signup">
                            Sign Up
                        </Link>
                    </li>
                    <li>
                        <Link className="nav-link" to="/hospital-signup">
                            Hospital Sign Up
                        </Link>
                    </li>
                    <li>
                        <Link className="nav-link" to="/hospital-login">
                            Hospital Login
                        </Link>
                    </li>
                </ul>
            </nav>
        </div>
    );
};

export default NavigationLog;
