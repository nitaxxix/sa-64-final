import { PatientInterface } from "./IPantient";
import { RemedyInterface } from "./IRemedy";
import { UserInterface } from "./IUser";

export interface AppointInterface {
  ID: string,
  AppointTime: Date,
  Todo: string,
  PatientID: number,
  Patient: PatientInterface,
  RemedyTypeID: number,
  RemedyType: RemedyInterface,
  UserDentistID: number,
  UserDentist: UserInterface,
}