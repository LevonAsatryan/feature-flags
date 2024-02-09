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
    return this.featureFlagRepository.save(featureFlag);
  }

  findAll() {
    return `This action returns all featureFlags`;
  }

  findOne(id: number) {
    return `This action returns a #${id} featureFlag`;
  }

  update(id: number, updateFeatureFlagDto: UpdateFeatureFlagDto) {
    return `This action updates a #${id} featureFlag`;
  }

  remove(id: number) {
    return `This action removes a #${id} featureFlag`;
  }
}
