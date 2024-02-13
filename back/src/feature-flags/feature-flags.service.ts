import { Injectable } from '@nestjs/common';
import { CreateFeatureFlagDto } from './dto/create-feature-flag.dto';
import { UpdateFeatureFlagDto } from './dto/update-feature-flag.dto';
import { FeatureFlag } from './entities/feature-flag.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

@Injectable()
export class FeatureFlagsService {
  constructor(
    @InjectRepository(FeatureFlag)
    private featureFlagRepository: Repository<FeatureFlag>
  ) {}

  create(createFeatureFlagDto: CreateFeatureFlagDto) {
    const featureFlag = new FeatureFlag();
    featureFlag.name = createFeatureFlagDto.name;
    featureFlag.companyId = createFeatureFlagDto.companyId;
    featureFlag.isEnabled = true;
    featureFlag.containerId = createFeatureFlagDto.containerId;
    return this.featureFlagRepository.save(featureFlag);
  }

  async findFeatureFlagsByContainerId(containerId: number) {
    const ff = await this.featureFlagRepository.query(
      `SELECT * FROM feature_flag WHERE container_id = ${containerId}`
    );

    return ff;
  }

  createTestFeatureForRoot(companyId: number, containerId: number) {
    const featureFlag = new FeatureFlag();
    featureFlag.name = 'test';
    featureFlag.companyId = companyId;
    featureFlag.isEnabled = true;
    featureFlag.containerId = containerId;
    return this.featureFlagRepository.save(featureFlag);
  }

  findAll() {
    return this.featureFlagRepository.find();
  }

  findOne(id: number) {
    return this.featureFlagRepository.findOne({ where: { id } });
  }

  async update(id: number, updateFeatureFlagDto: UpdateFeatureFlagDto) {
    const ff = await this.featureFlagRepository.findOne({ where: { id } });

    ff.containerId = updateFeatureFlagDto.containerId || ff.containerId;
    ff.name = updateFeatureFlagDto.name || ff.name;
    ff.isEnabled = updateFeatureFlagDto.isEnabled || ff.isEnabled;

    return this.featureFlagRepository.save(ff);
  }

  async remove(id: number) {
    const ff = await this.featureFlagRepository.findOne({ where: { id } });
    this.featureFlagRepository.delete({ id });
    return ff;
  }
}
