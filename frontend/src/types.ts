export type CPUMessage = {
  type: 'cpu';
  payload: {
    temprature: number;
  };
  timestamp: number;
};
