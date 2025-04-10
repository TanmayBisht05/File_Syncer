// frontend/pages/ProfilePage.jsx
import React from "react";
import { useAuthStore } from "../store/useAuthStore";
import "./ProfilePage.css";

const ProfilePage = () => {
  const { authUser } = useAuthStore();

  if (!authUser) {
    return <div className="profile-container">Not logged in</div>;
  }

  return (
    <div className="profile-container">
      <h2>Your Profile</h2>
      <div className="profile-info">
        <p><strong>Username:</strong> {authUser.username}</p>
        <p><strong>Email:</strong> {authUser.email}</p>
      </div>
    </div>
  );
};

export default ProfilePage;
