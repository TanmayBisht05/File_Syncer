import React from "react";
import { Link } from "react-router-dom";
import "./Navbar.css";
import { useAuthStore } from "../store/useAuthStore";

const Navbar = () => {
  const { authUser, logout } = useAuthStore();

  return (
    <nav className="navbar">
      <div className="navbar-left">
        <Link to="/" className="navbar-logo">
          FileSyncer
        </Link>
      </div>
      <div className="navbar-right">
        {authUser ? (
          <>
            <span className="navbar-user">Hi, {authUser.username}</span>
            <Link to="/profile" className="navbar-link">
              Profile
            </Link>
            <Link to="/" className="navbar-link">
              Home
            </Link>
            <button onClick={logout} className="navbar-logout">
              Logout
            </button>
          </>
        ) : (
          <>
            <Link to="/login" className="navbar-link">
              Login
            </Link>
            <Link to="/signup" className="navbar-link">
              Sign Up
            </Link>
          </>
        )}
      </div>
    </nav>
  );
};

export default Navbar;
