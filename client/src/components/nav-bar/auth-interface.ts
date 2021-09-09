import {User} from "../../user";
import React from "react";

export interface AuthInterface {
    currentUser?: User | undefined;
    setCurrentUser?:  React.Dispatch<React.SetStateAction<User | undefined>>
}
