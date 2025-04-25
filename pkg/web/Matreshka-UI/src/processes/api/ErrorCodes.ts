import { ToastMessageOptions, ToastServiceMethods } from "primevue";

export enum GrpcCodes {
  OK,
  CANCELLED,
  UNKNOWN,
  INVALID_ARGUMENT,
  DEADLINE_EXCEEDED,
  NOT_FOUND,
  ALREADY_EXISTS,
  PERMISSION_DENIED,
  RESOURCE_EXHAUSTED,
  FAILED_PRECONDITION,
  ABORTED,
  OUT_OF_RANGE,
  UNIMPLEMENTED,
  INTERNAL,
  UNAVAILABLE,
  DATA_LOSS,
  UNAUTHENTICATED,
}

export type GrpcError = {
  code: number;
  message: string;
  metadata: object;
};

export default function handleGrpcError(
  toastApi: ToastServiceMethods
): (err: GrpcError) => undefined {
  return (err: GrpcError) => {
    console.debug("got error", err);
    const msg: ToastMessageOptions = {
      closable: true,
      life: 5_000,
      summary: err.message,
      severity: "warn",
    };

    switch (err.code) {
      case (GrpcCodes.UNAVAILABLE, undefined):
        msg.severity = "error";
        break;
    }

    toastApi.add(msg);

    return;
  };
}
