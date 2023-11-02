import { useEffect, useState } from "react";
import { api } from "../../api";
import {
  Config as ConfigModel,
  NewConfig as NewConfigModel,
} from "../../models/Config";

async function createNewConfig(event: any, config: NewConfigModel) {
  event.preventDefault();

  try {
    await api.post("/config", config);
  } catch (e: any) {
    console.error(e.message);
  }
}

async function deleteConfig(e: any, id: number) {
  e.preventDefault();

  try {
    await api.delete(`/config/${id}`);
  } catch (e: any) {
    console.error(e.message);
  }
}

const Config = () => {
  const [config, setConfig] = useState<ConfigModel[]>();
  const [newConfigName, setNewConfigName] = useState<string>("");
  const [newConfigValue, setNewConfigValue] = useState<string>("");

  useEffect(() => {
    const getConfig = async () => {
      try {
        const config = await api.get("/config");
        setConfig(config.data.data);
      } catch (e: any) {
        console.error(e.message);
      }
    };

    getConfig();
  }, []);
  return (
    <div>
      <h2>Config</h2>
      <div>
        <h3>Create config</h3>
        <label htmlFor="name">Name</label>
        <input
          id="name"
          type="text"
          value={newConfigName}
          onChange={(e) => setNewConfigName(e.target.value)}
        />

        <label htmlFor="value">Value</label>
        <input
          id="value"
          type="text"
          value={newConfigValue}
          onChange={(e) => setNewConfigValue(e.target.value)}
        />
        <button
          onClick={(e) =>
            createNewConfig(e, { name: newConfigName, value: newConfigValue })
          }
        >
          Create
        </button>
      </div>

      <div>
        <h3>Manage config</h3>
        <ul>
          {/*  @ts-ignore*/}
          {config &&
            config.map((c: any, k) => {
              return (
                <li key={k}>
                  {c.name} - {c.value} -{" "}
                  <button onClick={(e) => deleteConfig(e, c.id)}>delete</button>
                </li>
              );
            })}
        </ul>
      </div>
    </div>
  );
};

export default Config;
