import { Injectable } from '@nestjs/common';
import { CreateCompanyDto } from './dto/create-company.dto';
import { UpdateCompanyDto } from './dto/update-company.dto';
import { Repository } from 'typeorm';
import { Company } from './entities/company.entity';
import { ContainersService } from 'src/containers/containers.service';
import { InjectRepository } from '@nestjs/typeorm';
import { FeatureFlagsService } from 'src/feature-flags/feature-flags.service';
import { UsersService } from 'src/users/users.service';

@Injectable()
export class CompanyService {
  constructor(
    @InjectRepository(Company)
    private readonly companyRepository: Repository<Company>,
    private readonly containerService: ContainersService,
    private readonly featureService: FeatureFlagsService,
    private readonly usersService: UsersService
  ) {}

  async create(createCompanyDto: CreateCompanyDto) {
    const company = new Company();
    company.name = createCompanyDto.name;
    const companyObj = await this.companyRepository.save(company);
    const rootContainer = await this.containerService.createRootContainer(
      companyObj.id
    );
    await this.featureService.createTestFeatureForRoot(
      companyObj.id,
      rootContainer.id
    );
    await this.usersService.createDefaultAdminUser(
      createCompanyDto.email,
      companyObj
    );
    return companyObj;
  }

  findAll() {
    return this.companyRepository.find();
  }

  findOne(id: number) {
    return this.companyRepository.findOneBy({ id });
  }

  async update(id: number, updateCompanyDto: UpdateCompanyDto) {
    await this.companyRepository.update(id, updateCompanyDto);
    return this.companyRepository.findOneBy({ id });
  }

  remove(id: number) {
    const company = this.companyRepository.findOneBy({ id });
    this.companyRepository.delete({ id });
    return company;
  }
}
