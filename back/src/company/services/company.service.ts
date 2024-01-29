import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Company } from '../schemas/company.schema';
import { Model } from 'mongoose';
import { CreateCompanyDTO } from '../DTO/company.dto';

@Injectable()
export class CompanyService {
  constructor(
    @InjectModel(Company.name) private companyModel: Model<Company>
  ) {}

  async create(company: CreateCompanyDTO): Promise<Company> {
    const createdCompany = new this.companyModel({
      name: company.name,
      featureFlags: [],
    });

    return createdCompany.save();
  }

  async findAll(): Promise<Company[]> {
    return this.companyModel.find().exec();
  }
}
