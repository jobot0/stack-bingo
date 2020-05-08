export interface DumbRepository {
  save(dumbName: string): Promise<void>;
}
