export class RingBuffer<T> {
  private buf: (T | undefined)[];
  private head = 0;
  private _size = 0;
  public readonly capacity: number;

  constructor(capacity: number) {
    this.capacity = capacity;
    this.buf = new Array(capacity);
  }

  push(item: T): void {
    this.buf[this.head] = item;
    this.head = (this.head + 1) % this.capacity;
    if (this._size < this.capacity) this._size++;
  }

  get size(): number {
    return this._size;
  }

  // Возвращает точки от самой старой к самой новой
  toArray(): T[] {
    const out: T[] = new Array(this._size);
    const start = (this.head - this._size + this.capacity) % this.capacity;
    for (let i = 0; i < this._size; i++) {
      out[i] = this.buf[(start + i) % this.capacity] as T;
    }
    return out;
  }
}
