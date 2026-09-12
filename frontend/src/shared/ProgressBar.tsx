interface IProgress {
  progress: number;
  label?: string;
  color?: string;
  showValue?: boolean;
}

function ProgressBar({ progress, label, showValue = false, color = '#02555e' }: IProgress) {
  return (
    <div style={{ display: 'flex', gap: '10px', alignItems: 'center' }}>
      <span>{label}</span>
      <div
        style={{
          width: 300,
          height: 20,
          background: '#eee',
          borderRadius: 10,
          position: 'relative',
        }}
      >
        {showValue && (
          <div
            style={{
              position: 'absolute',
              left: '50%',
              transform: 'translate(-50%)',
            }}
          >
            {progress} %
          </div>
        )}
        <div
          style={{
            width: `${progress}%`,
            height: '100%',
            background: color,
            borderRadius: 10,
            transition: 'width 0.3s',
          }}
        />
      </div>
    </div>
  );
}

export default ProgressBar;
