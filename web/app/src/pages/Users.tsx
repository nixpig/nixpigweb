import { useEffect, useState } from "react";
import { http } from "../services";

export const Users = () => {
  const [users, setUsers] = useState<any[]>([]);

  const [editingUserId, setEditingUserId] = useState<number>(-1);

  useEffect(() => {
    http.get("/api/user").then((res) => {
      setUsers(res.data.data);
    });
  }, []);

  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Username</th>
            <th>Email</th>
            <th>Admin</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {users ? (
            users.map((user) => {
              return (
                <tr key={user.id}>
                  <td>{user.id}</td>
                  <td>{user.username}</td>
                  <td>{user.email}</td>
                  <td>{JSON.stringify(user.is_admin)}</td>
                  <td>
                    {editingUserId !== user.id ? (
                      <button onClick={(e) => setEditingUserId(user.id)}>
                        Edit
                      </button>
                    ) : (
                      <div>
                        <button>Save</button> | <button>Cancel</button>
                      </div>
                    )}
                  </td>
                </tr>
              );
            })
          ) : (
            <tr>
              <td colSpan={5}>No users</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};
