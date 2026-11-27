type Listener = () => void;

export class SessionEvents {
  private listeners = new Set<Listener>();

  onSessionExpired(listener: Listener): () => void {
    this.listeners.add(listener);
    return () => this.listeners.delete(listener);
  }

  emitSessionExpired(): void {
    this.listeners.forEach((listener) => listener());
  }
}
