import CONFIG from "../config";
import { User } from "../types";

export const fetchUser = async (): Promise<User | null> => {
  try {
    const response = await fetch(`${CONFIG.backendURL}/api/user/whoami`, {
      method: "GET",
      credentials: "include",
    });

    if (response.ok) {
      const data = await response.json();
      return { userName: data.email };
    } else {
      return { userName: "" };
    }
  } catch (error: any) {
    console.error(error);
  } finally {
    return null;
  }
};
