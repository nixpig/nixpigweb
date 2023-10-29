import { useEffect, useState } from "react";
import { api } from "../api";

async function createNewConfig(e: any, name: string, value: string) {
  e.preventDefault();

  const res = await api.post("/config", {
    name,
    value,
  });
  console.log({ res });
}
const Config = () => {
  const [config, setConfig] = useState<{ data: any[] }>({ data: [] });
  const [newConfigName, setNewConfigName] = useState<string>("");
  const [newConfigValue, setNewConfigValue] = useState<string>("");

  useEffect(() => {
    const getConfig = async () => {
      const config = await api.get("/config");

      setConfig(config.data);
      console.log(config.data);
    };

    getConfig();
  }, []);
  return (
    <div>
      <h2>Config</h2>
      <div>
        <h3>Create</h3>
        <input
          type="text"
          value={newConfigName}
          onChange={(e) => setNewConfigName(e.target.value)}
        />
        <input
          type="text"
          value={newConfigValue}
          onChange={(e) => setNewConfigValue(e.target.value)}
        />
        <button
          onClick={(e) => createNewConfig(e, newConfigName, newConfigValue)}
        >
          Create
        </button>
      </div>

      <div>
        <h3>View</h3>
        <ul>
          {/*  @ts-ignore*/}
          {config.data &&
            config.data.map((c: any, k) => {
              return (
                <li key={k}>
                  {c.name} - {c.value}
                </li>
              );
            })}
        </ul>
      </div>
    </div>
  );
};

export default Config;
