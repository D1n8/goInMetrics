import { useEffect, useState } from 'react';
import ProgressBar from './shared/ProgressBar';

interface IMetrics {
  cpu: {
    model: string;
    load_percent: number;
    temperature: number;
    frequency: {
      max_mhz: number;
      work_mhz: number;
    }
  };
  mem: {
    total: number;
    used: number;
    used_percent: number;
    avail: number;
    swap_total: number;
    swap_used: number;
  }
}

function App() {
  const [metrics, setMetrics] = useState<IMetrics | null>();

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/ws');
    
    ws.onmessage = (event) => {
      setMetrics(JSON.parse(event.data));
    };

    return () => ws.close();
  }, []);

  return (
    <div>
      <h2>{metrics?.cpu.model}</h2>
      <ProgressBar progress={metrics?.mem.used_percent.toFixed(2) ?? 0} label='Памяти используется' showValue/>
      <ProgressBar progress={metrics?.cpu.load_percent.toFixed(2) ?? 0} label='Проциссер' color='#40c97e' showValue/>
    </div>
  );
}

export default App;