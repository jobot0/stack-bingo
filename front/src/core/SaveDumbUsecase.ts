import { DumbRepository } from "./DumbRepository";
export class SaveDumbUsecase {
  dumbRepository: DumbRepository;

  constructor(dumbRepository: DumbRepository) {
    this.dumbRepository = dumbRepository;
  }

  async execute(dumbName: string): Promise<void> {
    return this.dumbRepository.save(dumbName);
  }
}
