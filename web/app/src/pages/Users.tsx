import axios from "axios";
import { useEffect, useState } from "react";

export const Users = () => {
  const [users, setUsers] = useState<any[]>([]);

  useEffect(() => {
    axios.get("http://localhost:8081/api/user").then((res) => {
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
                </tr>
              );
            })
          ) : (
            <tr>
              <td colSpan={4}>No users</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};
