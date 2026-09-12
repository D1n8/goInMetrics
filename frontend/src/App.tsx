import ProgressBar from './ProgressBar';

function App() {
  return (
    <div
      style={{
        height: '100%',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
      }}
    >
      <div>
        <ProgressBar label="Usage" progress={14} showValue />
      </div>
    </div>
  );
}

export default App;
